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

func (s *Service) GetUserByEmail(email string) (*User, bool) {
	return s.repo.fetchUserByEmail(email)
}

func (s *Service) Signup(user *User, password string) (*SignupResult, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	user, ok := s.createUser(user, hashedPassword)
	if !ok {
		return nil, errors.New("failed to create user")
	}
	jwt, err := s.createJwt(user)
	if err != nil {
		return nil, err
	}
	refreshJWT, err := s.createRefreshJwt(user)
	if err != nil {
		return nil, err
	}

	return &SignupResult{
		User:                  user,
		AccessToken:           jwt,
		ExpiresAt:             "1",
		RefreshToken:          refreshJWT,
		RefreshTokenExpiresAt: "3",
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
