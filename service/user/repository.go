package user

import "github.com/katerji/ecommerce/db"

type repo struct{}

func (repo) fetchUserByEmail(email string) (*User, bool) {
	user, ok := db.FetchOne[dbUser, User](fetchUserByEmailQuery, email)
	return &user, ok
}

func (repo) fetchUserByPhoneNumber(phoneNumber string) (*User, bool) {
	user, ok := db.FetchOne[dbUser, User](fetchUserByPhoneNumberQuery, phoneNumber)
	return &user, ok
}

func (repo) insertUser(user *User, password string) (*User, bool) {
	userID, ok := db.Insert(insertUserQuery, user.Name, user.Email, user.PhoneNumber, password)
	user.ID = userID

	return user, ok
}

func (repo) insertAddress(address Address) (Address, bool) {
	addressID, ok := db.Insert(insertAddressQuery, address.UserID, address.AddressLine1, address.AddressLine2, address.Country, address.City, address.State, address.ZipCode)
	address.ID = addressID

	return address, ok
}

func (repo) updateAddress(address Address) bool {
	return db.Update(updateAddressQuery, address.AddressLine1, address.AddressLine2, address.Country, address.City, address.State, address.ZipCode, address.ID)
}

func (repo) deleteAddress(addressID int) bool {
	return db.Delete(deleteAddressQuery, addressID)
}

func (repo) fetchAddresses(userID int) map[int]Address {
	addresses := db.Fetch[dbAddress, Address](fetchAddressesQuery, userID)

	addressMap := make(map[int]Address, len(addresses))
	for _, a := range addresses {
		addressMap[a.ID] = a
	}

	return addressMap
}
