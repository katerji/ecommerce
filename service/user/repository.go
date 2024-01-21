package user

import "github.com/katerji/ecommerce/db"

type repo struct{}

func (repo) fetchUserByEmail(email string) (User, bool) {
	return db.FetchOne[dbUser, User](fetchUserByEmailQuery, email)
}

func (repo) fetchUserByPhoneNumber(phoneNumber string) (User, bool) {
	return db.FetchOne[dbUser, User](fetchUserByPhoneNumberQuery, phoneNumber)
}

func (repo) insertUser(user User, password string) (User, bool) {
	userID, ok := db.Insert(insertUserQuery, user.Name, user.Email, user.PhoneNumber, password)
	user.ID = userID

	return user, ok
}
