package user

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func randomUser() User {
	return User{
		Email:       fmt.Sprintf("%s@gmail.com", uuid.NewString()[0:8]),
		Name:        fmt.Sprintf("%s", uuid.NewString()[0:8]),
		PhoneNumber: fmt.Sprintf("%s", uuid.NewString()[0:4]),
	}
}

func TestInsert(t *testing.T) {
	user := randomUser()
	newUser, ok := repo{}.insertUser(user, "pass")
	if !ok {
		t.Fatal("failed to insert user")
	}

	if newUser.Email != user.Email {
		t.Fatalf("expected email %s, got %s", user.Email, newUser.Email)
	}

	if newUser.PhoneNumber != user.PhoneNumber {
		t.Fatalf("expected email %s, got %s", user.Email, newUser.Email)
	}

	if newUser.Name != user.Name {
		t.Fatalf("expected email %s, got %s", user.Email, newUser.Email)
	}
}

func TestSelectByEmail(t *testing.T) {
	user := randomUser()
	r := repo{}
	r.insertUser(user, "pass")

	fetchedUser, ok := r.fetchUserByEmail(user.Email)
	if !ok {
		t.Fatalf("failed to fetch user by email")
	}

	if fetchedUser.Email != user.Email {
		t.Fatalf("expected email %s, got %s", user.Email, fetchedUser.Email)
	}

	if fetchedUser.PhoneNumber != user.PhoneNumber {
		t.Fatalf("expected email %s, got %s", user.Email, fetchedUser.Email)
	}

	if fetchedUser.Name != user.Name {
		t.Fatalf("expected email %s, got %s", user.Email, fetchedUser.Email)
	}

}

func TestSelectByPhoneNumber(t *testing.T) {
	user := randomUser()
	r := repo{}
	r.insertUser(user, "pass")

	fetchedUser, ok := r.fetchUserByPhoneNumber(user.PhoneNumber)
	if !ok {
		t.Fatalf("failed to fetch user by email")
	}

	if fetchedUser.Email != user.Email {
		t.Fatalf("expected email %s, got %s", user.Email, fetchedUser.Email)
	}

	if fetchedUser.PhoneNumber != user.PhoneNumber {
		t.Fatalf("expected email %s, got %s", user.Email, fetchedUser.Email)
	}

	if fetchedUser.Name != user.Name {
		t.Fatalf("expected email %s, got %s", user.Email, fetchedUser.Email)
	}

}
