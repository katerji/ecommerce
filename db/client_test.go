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

func TestFetch(t *testing.T) {
	users := Fetch[dbUser, user]("SELECT id, email FROM user LIMIT 50")
	for _, u := range users {
		fmt.Println(u.ID)
		fmt.Println(u.Name)
		fmt.Println(u.Email)
	}
}

func TestFetchOne(t *testing.T) {
	u := FetchOne[dbUser, user]("SELECT id, email FROM user LIMIT 50")
	fmt.Println(u.ID)
	fmt.Println(u.Name)
	fmt.Println(u.Email)
}
