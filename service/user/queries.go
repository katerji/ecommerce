package user

const (
	fetchUserByEmailQuery       = "SELECT id, name, email, phone_number FROM user WHERE email = ?"
	fetchUserByPhoneNumberQuery = "SELECT id, name, email, phone_number FROM user WHERE phone_number = ?"
	insertUserQuery             = "INSERT INTO user (name, email, phone_number, password) VALUES (?, ?, ?, ?)"
)
