package user

const (
	fetchUserByEmailQuery       = "SELECT id, name, email, phone_number FROM user WHERE email = ?"
	fetchUserByPhoneNumberQuery = "SELECT id, name, email, phone_number FROM user WHERE phone_number = ?"
	insertUserQuery             = "INSERT INTO user (name, email, phone_number, password) VALUES (?, ?, ?, ?)"
)

const (
	fetchAddressesQuery = "SELECT id, user_id, address_line1, address_line2, city, state, zip_code, country FROM address WHERE user_id = ? ORDER BY created_on DESC"
	insertAddressQuery  = "INSERT INTO address (user_id, address_line1, address_line2, country, city, state, zip_code) VALUES (?, ?, ?, ?, ?, ?, ?)"
	updateAddressQuery  = "UPDATE address SET address_line1 = ?, address_line2 = ?, country = ?, city = ?, state = ?, zip_code = ? WHERE id = ?"
	deleteAddressQuery  = "DELETE FROM address WHERE id = ?"
)
