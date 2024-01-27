package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/katerji/ecommerce/envs"
	_ "strconv"
	_ "testing"
	"time"
)

var ErrNoRows = errors.New("no rows found")

type Client struct {
	*sqlx.DB
}

var instance *Client

func getDbInstance() *Client {
	if instance == nil {
		instance, _ = getDbClient()
	}
	return instance
}

func getDbClient() (*Client, error) {
	dbHost := envs.GetInstance().GetDbHost()
	dbUser := envs.GetInstance().GetDbUser()
	dbPort := envs.GetInstance().GetDbPort()
	dbPass := envs.GetInstance().GetDbPassword()
	dbName := envs.GetInstance().GetDbName()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &Client{
		db,
	}, nil
}

func Fetch[T Reader, W any](query string, args ...any) ([]W, error) {
	client := getDbInstance()

	var dbModels []T

	err := client.Select(&dbModels, query, args...)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(dbModels) == 0 {
		return nil, ErrNoRows
	}
	var returns []W
	for _, m := range dbModels {
		returns = append(returns, m.ToModel().(W))
	}

	return returns, nil
}

func FetchOne[T Reader, W any](query string, args ...any) (W, error) {
	client := getDbInstance()

	var dbModel T
	err := client.Get(&dbModel, query, args...)
	if !errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err)
		return dbModel.ToModel().(W), ErrNoRows
	}
	return dbModel.ToModel().(W), nil
}

func Insert(query string, args ...any) (int, error) {
	client := getDbInstance()
	prepare, err := client.Prepare(query)
	if err != nil {
		fmt.Printf("err inserting: %v", err)
		return 0, err
	}
	result, err := prepare.Exec(args...)
	if err != nil {
		fmt.Printf("err inserting: %v", err)
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("err inserting: %v", err)
		return 0, err
	}

	return int(insertID), nil
}

func Update(query string, args ...any) error {
	client := getDbInstance()
	prepare, err := client.Prepare(query)
	if err != nil {
		fmt.Printf("err updating: %v", err)
		return err
	}
	result, err := prepare.Exec(args...)
	if err != nil {
		fmt.Printf("err updating: %v", err)
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		fmt.Printf("err updating: %v", err)
		return err
	}

	return nil
}

func Delete(query string, args ...any) error {
	client := getDbInstance()
	prepare, err := client.Prepare(query)
	if err != nil {
		fmt.Printf("err updating: %v", err)
		return err
	}
	result, err := prepare.Exec(args...)
	if err != nil {
		fmt.Printf("err updating: %v", err)
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		fmt.Printf("err updating: %v", err)
		return err
	}

	return nil
}
