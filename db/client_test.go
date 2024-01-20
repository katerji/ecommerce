package db

import "testing"

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
