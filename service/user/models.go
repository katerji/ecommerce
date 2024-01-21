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

type Address struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	Country      string `json:"country"`
}

type dbAddress struct {
	ID           uint   `db:"id"`
	UserID       uint   `db:"user_id"`
	AddressLine1 string `db:"address_line1"`
	AddressLine2 string `db:"address_line2"`
	City         string `db:"city"`
	State        string `db:"state"`
	ZipCode      string `db:"zip_code"`
	Country      string `db:"country"`
}

func (a dbAddress) ToModel() any {
	return Address{
		ID:           int(a.ID),
		UserID:       int(a.UserID),
		AddressLine1: a.AddressLine1,
		AddressLine2: a.AddressLine2,
		City:         a.City,
		State:        a.State,
		ZipCode:      a.ZipCode,
		Country:      a.Country,
	}
}

type SignupResult struct {
	User                  *User
	AccessToken           string
	ExpiresAt             string
	RefreshToken          string
	RefreshTokenExpiresAt string
}
