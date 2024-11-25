package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

func ConnectPostgres(connString string) (*DB, error) {
	d, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	err = testDB(d)
	if err != nil {
		return nil, err
	}

	dbConn.DB = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("*** Pinged database successfully! ***")
	return nil
}
