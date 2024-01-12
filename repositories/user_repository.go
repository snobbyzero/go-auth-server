package repositories

import (
	"context"
	"go_auth_server/database"
	"go_auth_server/database/models"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(ctx context.Context) *UserRepository {
	pool, err := database.GetDB(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return &UserRepository{pool}
}

func (userRepository *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.UserModel, error) {
	var user models.UserModel
	err := userRepository.pool.QueryRow(ctx, "SELECT id, password FROM users WHERE email=$1", email).Scan(&user.ID, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) CreateUser(ctx context.Context, email, username, password string) (string, error) {
	var id int
	t := time.Now()
	err := userRepository.pool.QueryRow(ctx, "INSERT INTO users (email, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING ID", email, username, password, t, t).Scan(&id)

	return strconv.Itoa(id), err
}
