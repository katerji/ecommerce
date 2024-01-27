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
	repository *repository
}

func New() *Service {
	return &Service{
		repository: &repository{},
	}
}

func (s *Service) getUserByEmail(email string) (*UserWithPass, error) {
	return s.repository.fetchUserByEmail(email)
}

func (s *Service) getUserByPhoneNumber(phoneNumber string) (*UserWithPass, error) {
	return s.repository.fetchUserByPhoneNumber(phoneNumber)
}

func (s *Service) LoginWithEmail(email string, password string) (*LoginResult, error) {
	userWithPass, err := s.getUserByEmail(email)
	if err != nil || userWithPass == nil {
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
	userWithPass, err := s.getUserByPhoneNumber(phoneNumber)
	if err != nil || userWithPass == nil {
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
	user, err = s.createUser(user, hashedPassword)
	if err != nil {
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

func (s *Service) createUser(user *User, password string) (*User, error) {
	return s.repository.insertUser(user, password)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Service) GetAddresses(userID int) (map[int]Address, error) {
	return s.repository.fetchAddresses(userID)
}

func (s *Service) CreateAddress(address *Address) (*Address, error) {
	return s.repository.insertAddress(address)
}

func (s *Service) UpdateAddress(address *Address) error {
	return s.repository.updateAddress(address)
}

func (s *Service) DeleteAddress(addressID int) error {
	return s.repository.deleteAddress(addressID)
}
