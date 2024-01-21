package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *repo
}

func New() *Service {
	return &Service{
		repo: &repo{},
	}
}

func (s *Service) getUserByEmail(email string) (*User, bool) {
	return s.repo.fetchUserByEmail(email)
}

func (s *Service) getUserByPhoneNumber(phoneNumber string) (*User, bool) {
	return s.repo.fetchUserByPhoneNumber(phoneNumber)
}

func (s *Service) LoginWithEmail(email string) (*LoginResult, error) {
	user, ok := s.getUserByEmail(email)
	if !ok || user == nil {

	}

	pair, err := s.createJwt(user)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		User:    user,
		jwtPair: pair,
	}, nil
}

func (s *Service) LoginWithPhoneNumber(phoneNumber string) (*LoginResult, error) {
	user, ok := s.getUserByPhoneNumber(phoneNumber)
	if !ok || user == nil {
		return nil, errors.New("account not found")
	}

	pair, err := s.createJwt(user)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		User:    user,
		jwtPair: pair,
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
		jwtPair: jwtPair,
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
