package user

type User struct {
	ID          int
	Email       string
	Name        string
	PhoneNumber string
}

type dbUser struct {
	ID          uint   `db:"id"`
	Email       string `db:"email"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
}

func (u dbUser) ToModel() any {
	return User{
		ID:          int(u.ID),
		Email:       u.Email,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}
}
