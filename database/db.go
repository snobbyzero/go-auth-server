package database

import (
	"database/sql"
	"fmt"
	"go_auth_server/config"
)


var db *sql.DB

func GetDB() (*sql.DB, error) {
	var err error
	if db == nil {
		db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Dbname)) // create db_config.go file with these constants
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
