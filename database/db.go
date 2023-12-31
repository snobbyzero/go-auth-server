package database

import (
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"go_auth_server/config"
	"os"
)


var conn *pgx.Conn

func GetDB() (*pgx.Conn, error) {
	var err error
	if conn == nil {
		conn, err = pgx.Connect(pgx.ConnConfig{
			Host: config.Host,
			Port: config.Port,
			User: config.User,
			Password: config.Password,
			Database: config.Dbname,
		}) // create config/db_config.go file with these constants
		if err != nil {
			return nil, err
		}
	}
	return conn, nil
}

func CreateTables() error {
	conn, err := GetDB()
	query, err := os.ReadFile("./config/table.sql")
	if err != nil {
		return err
	}
	_, err = conn.Exec(string(query))
	if err != nil {
		return err
	}
	return nil
}