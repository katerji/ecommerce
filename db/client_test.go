package db

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	c, err := getDbClient()
	if err != nil {
		t.Fatal(err)
	}
	err = c.Ping()
	if err != nil {
		t.Fatal(err)
	}

}

type dbUser struct {
	ID    uint   `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type user struct {
	ID    uint
	Name  string
	Email string
}

func (u dbUser) ToModel() any {
	return user{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}

}
func TestFetch(t *testing.T) {
	users, err := Fetch[dbUser, user]("SELECT id, email FROM user LIMIT 50")
	if err != nil {
		t.Fatalf("failed to fetch users with err: %v", err)
	}
	for _, u := range users {
		fmt.Println(u.ID)
		fmt.Println(u.Name)
		fmt.Println(u.Email)
	}
}

func TestFetchOne(t *testing.T) {
	u, ok := FetchOne[dbUser, user]("SELECT id, email FROM user LIMIT 50")
	if !ok {
		t.Fatal("failed to fetch users")
	}
	fmt.Println(u.ID)
	fmt.Println(u.Name)
	fmt.Println(u.Email)
}
