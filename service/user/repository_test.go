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

func randomAddress(userID int) *Address {
	return &Address{
		UserID:       userID,
		AddressLine1: fmt.Sprintf("%s", uuid.NewString()[0:4]),
		AddressLine2: fmt.Sprintf("%s", uuid.NewString()[0:4]),
		City:         fmt.Sprintf("%s", uuid.NewString()[0:4]),
		State:        fmt.Sprintf("%s", uuid.NewString()[0:4]),
		ZipCode:      fmt.Sprintf("%s", uuid.NewString()[0:4]),
		Country:      fmt.Sprintf("%s", uuid.NewString()[0:4]),
	}
}

func TestInsertUser(t *testing.T) {
	user := randomUser()
	newUser, ok := repo{}.insertUser(&user, "pass")
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
	r.insertUser(&user, "pass")

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
	r.insertUser(&user, "pass")

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

func Test_repo_insertAddress(t *testing.T) {
	tests := []struct {
		name            string
		expectedAddress *Address
		want            bool
	}{
		{
			"test insert 1",
			randomAddress(1),
			true,
		},
		{
			"test insert 2",
			randomAddress(2),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := repo{}
			ad, got := re.insertAddress(tt.expectedAddress)
			if got != tt.want {
				t.Errorf("insertAddress() = %v, want %v", got, tt.want)
			}
			allAddresses, err := re.fetchAddresses(tt.expectedAddress.UserID)
			if err != nil {
				t.Errorf("failed to fetch addresses with error: %v", err)
			}
			insertedAddress := allAddresses[ad.ID]
			if tt.expectedAddress.AddressLine1 != insertedAddress.AddressLine1 {
				t.Errorf("AddressLine1 mismatch: got %v, want %v", insertedAddress.AddressLine1, tt.expectedAddress.AddressLine1)
			}

			if tt.expectedAddress.UserID != insertedAddress.UserID {
				t.Errorf("UserID missmatch: got %v, want %v", insertedAddress.UserID, tt.expectedAddress.UserID)
			}
			if tt.expectedAddress.AddressLine2 != insertedAddress.AddressLine2 {
				t.Errorf("AddressLine2 mismatch: got %v, want %v", insertedAddress.AddressLine2, tt.expectedAddress.AddressLine2)
			}

			if tt.expectedAddress.City != insertedAddress.City {
				t.Errorf("City mismatch: got %v, want %v", insertedAddress.City, tt.expectedAddress.City)
			}

			if tt.expectedAddress.State != insertedAddress.State {
				t.Errorf("State mismatch: got %v, want %v", insertedAddress.State, tt.expectedAddress.State)
			}

			if tt.expectedAddress.ZipCode != insertedAddress.ZipCode {
				t.Errorf("ZipCode mismatch: got %v, want %v", insertedAddress.ZipCode, tt.expectedAddress.ZipCode)
			}

			if tt.expectedAddress.Country != insertedAddress.Country {
				t.Errorf("Country mismatch: got %v, want %v", insertedAddress.Country, tt.expectedAddress.Country)
			}
		})
	}
}

func Test_repo_updateAddress(t *testing.T) {
	tests := []struct {
		name            string
		expectedAddress *Address
		want            bool
	}{
		{
			"test 1",
			&Address{
				ID:           1,
				UserID:       1,
				AddressLine1: "123",
				AddressLine2: "456",
				City:         "city",
				State:        "state",
				ZipCode:      "zip",
				Country:      "country",
			},
			true,
		},
		{
			"test 2",
			&Address{
				ID:           2,
				UserID:       1,
				AddressLine1: "new add",
				AddressLine2: "new add 2",
				City:         "city 2",
				State:        "state 2",
				ZipCode:      "zip 2",
				Country:      "country 2",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := repo{}
			if got := re.updateAddress(tt.expectedAddress); got != tt.want {
				t.Errorf("updateAddress() = %v, want %v", got, tt.want)
			}
			allAddresses, err := re.fetchAddresses(tt.expectedAddress.UserID)
			if err != nil {
				t.Fatalf("failed to fetch addresses with err: %v", err)
			}
			updatedAddress := allAddresses[tt.expectedAddress.ID]
			if tt.expectedAddress.AddressLine1 != updatedAddress.AddressLine1 {
				t.Errorf("AddressLine1 mismatch: got %v, want %v", updatedAddress.AddressLine1, tt.expectedAddress.AddressLine1)
			}

			// Check if other fields were updated correctly (similar checks for each field)
			if tt.expectedAddress.AddressLine2 != updatedAddress.AddressLine2 {
				t.Errorf("AddressLine2 mismatch: got %v, want %v", updatedAddress.AddressLine2, tt.expectedAddress.AddressLine2)
			}

			if tt.expectedAddress.City != updatedAddress.City {
				t.Errorf("City mismatch: got %v, want %v", updatedAddress.City, tt.expectedAddress.City)
			}

			if tt.expectedAddress.State != updatedAddress.State {
				t.Errorf("State mismatch: got %v, want %v", updatedAddress.State, tt.expectedAddress.State)
			}

			if tt.expectedAddress.ZipCode != updatedAddress.ZipCode {
				t.Errorf("ZipCode mismatch: got %v, want %v", updatedAddress.ZipCode, tt.expectedAddress.ZipCode)
			}

			if tt.expectedAddress.Country != updatedAddress.Country {
				t.Errorf("Country mismatch: got %v, want %v", updatedAddress.Country, tt.expectedAddress.Country)
			}
		})
	}
}

func Test_repo_deleteAddress(t *testing.T) {
	tests := []struct {
		name            string
		expectedAddress *Address
		want            bool
	}{
		{
			"test delete 1",
			randomAddress(1),
			true,
		},
		{
			"test delete 2",
			randomAddress(2),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := repo{}
			ad, _ := re.insertAddress(tt.expectedAddress)
			got := re.deleteAddress(ad.ID)
			if got != tt.want {
				t.Errorf("insertAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
