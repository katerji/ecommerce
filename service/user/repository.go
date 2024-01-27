package user

import (
	"github.com/katerji/ecommerce/db"
)

type repository struct{}

func (repository) fetchUserByEmail(email string) (*UserWithPass, error) {
	user, err := db.FetchOne[dbUser, UserWithPass](fetchUserByEmailQuery, email)
	return &user, err
}

func (repository) fetchUserByPhoneNumber(phoneNumber string) (*UserWithPass, error) {
	user, err := db.FetchOne[dbUser, UserWithPass](fetchUserByPhoneNumberQuery, phoneNumber)
	return &user, err
}

func (repository) insertUser(user *User, password string) (*User, error) {
	userID, err := db.Insert(insertUserQuery, user.Name, user.Email, user.PhoneNumber, password)
	user.ID = userID

	return user, err
}

func (repository) insertAddress(address *Address) (*Address, error) {
	addressID, err := db.Insert(insertAddressQuery, address.UserID, address.AddressLine1, address.AddressLine2, address.Country, address.City, address.State, address.ZipCode)
	address.ID = addressID

	return address, err
}

func (repository) updateAddress(address *Address) error {
	return db.Update(updateAddressQuery, address.AddressLine1, address.AddressLine2, address.Country, address.City, address.State, address.ZipCode, address.ID)
}

func (repository) deleteAddress(addressID int) error {
	return db.Delete(deleteAddressQuery, addressID)
}

func (repository) fetchAddresses(userID int) (map[int]Address, error) {
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
