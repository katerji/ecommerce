package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	emailNotFoundErr       = errors.New("email not found")
	phoneNumberNotFoundErr = errors.New("phone number not found")
	incorrectPasswordErr   = errors.New("incorrect password")
)

type Service struct {
	repo *repo
}

func (s *Service) InitService() {
	s.repo = &repo{}
}

func (s *Service) getUserByEmail(email string) (*UserWithPass, bool) {
	return s.repo.fetchUserByEmail(email)
}

func (s *Service) getUserByPhoneNumber(phoneNumber string) (*UserWithPass, bool) {
	return s.repo.fetchUserByPhoneNumber(phoneNumber)
}

func (s *Service) LoginWithEmail(email string, password string) (*LoginResult, error) {
	userWithPass, ok := s.getUserByEmail(email)
	if !ok || userWithPass == nil {
		return nil, emailNotFoundErr
	}

	if !validPassword(userWithPass.Password, password) {
		return nil, incorrectPasswordErr
	}
	user := userWithPass.User

	pair, err := s.createJwt(user)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		User:    user,
		JWTPair: pair,
	}, nil
}

func (s *Service) LoginWithPhoneNumber(phoneNumber string, password string) (*LoginResult, error) {
	userWithPass, ok := s.getUserByPhoneNumber(phoneNumber)
	if !ok || userWithPass == nil {
		return nil, phoneNumberNotFoundErr
	}

	if !validPassword(userWithPass.Password, password) {
		return nil, incorrectPasswordErr
	}
	user := userWithPass.User

	pair, err := s.createJwt(user)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		User:    user,
		JWTPair: pair,
	}, nil
}

func (s *Service) Signup(user *User, password string) (*LoginResult, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	user, ok := s.createUser(user, hashedPassword)
	if !ok {
		return nil, errors.New("failed to create user")
	}
	jwtPair, err := s.createJwt(user)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		User:    user,
		JWTPair: jwtPair,
	}, nil
}

func (s *Service) createUser(user *User, password string) (*User, bool) {
	return s.repo.insertUser(user, password)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
