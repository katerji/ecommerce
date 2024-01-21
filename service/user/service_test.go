package user

import (
	"testing"
)

func TestService_Signup(t *testing.T) {
	user := randomUser()
	password := "123456"

	s := New()

	_, err := s.Signup(&user, password)
	if err != nil {
		t.Fatalf("failed to sign up user with err: %v", err)
	}
}

func TestService_LoginWithEmail(t *testing.T) {
	user := randomUser()
	p := "11111"

	s := New()
	result, err := s.Signup(&user, p)
	if err != nil {
		t.Errorf("failed to sign up user with err: %v", err)
	}

	email := result.User.Email
	result, err = s.LoginWithEmail(email)
	if err != nil {
		t.Fatalf("failed to login with err: %v", err)
	}

	if user.PhoneNumber != result.User.PhoneNumber {
		t.Errorf("[%s] expected user phone number to be %s, got %s", "Email-Login:", user.PhoneNumber, result.User.PhoneNumber)
	}

	if user.PhoneNumber != result.User.PhoneNumber {
		t.Errorf("[%s] expected user name to be %s, got %s", "Email-Login:", user.Name, result.User.Name)
	}

	if user.Email != result.User.Email {
		t.Errorf("[%s] expected user email to be %s, got %s", "Email-Login:", user.Email, result.User.Email)
	}
}

func TestService_LoginWithPhone(t *testing.T) {
	user := randomUser()
	p := "11111"

	s := New()
	result, err := s.Signup(&user, p)
	if err != nil {
		t.Errorf("failed to sign up user with err: %v", err)
	}

	phoneNumber := result.User.PhoneNumber
	result, err = s.LoginWithPhoneNumber(phoneNumber)
	if err != nil {
		t.Fatalf("failed to login with err: %v", err)
	}

	if user.PhoneNumber != result.User.PhoneNumber {
		t.Errorf("[%s] expected user phone number to be %s, got %s", "Email-Login:", user.PhoneNumber, result.User.PhoneNumber)
	}

	if user.PhoneNumber != result.User.PhoneNumber {
		t.Errorf("[%s] expected user name to be %s, got %s", "Email-Login:", user.Name, result.User.Name)
	}

	if user.Email != result.User.Email {
		t.Errorf("[%s] expected user email to be %s, got %s", "Email-Login:", user.Email, result.User.Email)
	}
}

func TestService_VerifyJWT(t *testing.T) {
	user := randomUser()
	pass := "123"

	s := New()

	result, err := s.Signup(&user, pass)
	if err != nil {
		t.Fatalf("failed to signup user with err: %v", err)
	}
	_, err = s.verifyToken(result.jwtPair.AccessToken)
	if err != nil {
		t.Errorf("failed to verify access token with err: %v", err)
	}
	_, err = s.verifyRefreshToken(result.jwtPair.RefreshToken)
	if err != nil {
		t.Errorf("failed to verify refresh token with err: %v", err)
	}
}
