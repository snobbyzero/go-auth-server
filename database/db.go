package database

import (
	"context"
	"fmt"
	"go_auth_server/config"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var pool *pgxpool.Pool

var URL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?pool_max_conns=10", config.User, config.Password, config.Host, config.Port, config.Dbname)

func GetDB(ctx context.Context) (*pgxpool.Pool, error) {
	var err error
	if pool == nil {
		pool, err = connectDB(ctx)
		return pool, err
	}
	return pool, nil
}

func connectDB(ctx context.Context) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	return pgxpool.New(ctx, URL)
}

func CreateTables(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	pool, err := GetDB(ctx)
	query, err := os.ReadFile("./config/table.sql")
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, string(query))
	if err != nil {
		return err
	}
	return nil
}
