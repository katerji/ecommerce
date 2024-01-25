package user

import (
	"github.com/katerji/ecommerce/db"
)

type repo struct{}

func (repo) fetchUserByEmail(email string) (*UserWithPass, error) {
	user, err := db.FetchOne[dbUser, UserWithPass](fetchUserByEmailQuery, email)
	return &user, err
}

func (repo) fetchUserByPhoneNumber(phoneNumber string) (*UserWithPass, error) {
	user, err := db.FetchOne[dbUser, UserWithPass](fetchUserByPhoneNumberQuery, phoneNumber)
	return &user, err
}

func (repo) insertUser(user *User, password string) (*User, error) {
	userID, err := db.Insert(insertUserQuery, user.Name, user.Email, user.PhoneNumber, password)
	user.ID = userID

	return user, err
}

func (repo) insertAddress(address *Address) (*Address, error) {
	addressID, err := db.Insert(insertAddressQuery, address.UserID, address.AddressLine1, address.AddressLine2, address.Country, address.City, address.State, address.ZipCode)
	address.ID = addressID

	return address, err
}

func (repo) updateAddress(address *Address) bool {
	return db.Update(updateAddressQuery, address.AddressLine1, address.AddressLine2, address.Country, address.City, address.State, address.ZipCode, address.ID)
}

func (repo) deleteAddress(addressID int) bool {
	return db.Delete(deleteAddressQuery, addressID)
}

func (repo) fetchAddresses(userID int) (map[int]Address, error) {
	addresses, err := db.Fetch[dbAddress, Address](fetchAddressesQuery, userID)
	if err != nil {
		return nil, err
	}
	addressMap := make(map[int]Address, len(addresses))
	for _, a := range addresses {
		addressMap[a.ID] = a
	}

	return addressMap, nil
}
