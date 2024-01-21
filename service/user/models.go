package user

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type UserWithPass struct {
	*User
	Password string
}

type dbUser struct {
	ID          uint   `db:"id"`
	Email       string `db:"email"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
	Password    string `db:"password"`
}

func (u dbUser) ToModel() any {
	user := &User{
		ID:          int(u.ID),
		Email:       u.Email,
		Name:        u.Name,
		PhoneNumber: u.PhoneNumber,
	}
	return UserWithPass{
		User:     user,
		Password: u.Password,
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

type JWTPair struct {
	AccessToken      string `json:"access_token"`
	ExpiresAt        int64  `json:"expires_at"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresAt int64  `json:"refresh_expires_at"`
}

type LoginResult struct {
	User    *User
	JWTPair *JWTPair
}
