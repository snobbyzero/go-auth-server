package database

import "database/sql"

var db *sql.DB

func GetDB() (*sql.DB, error) {
	var err error
	if db == nil {
		db, err = sql.Open("postgres", "user=postgres password=postgres dbname=go-auth-server sslmode=disable")
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
