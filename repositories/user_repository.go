package repositories

import (
	"github.com/jackc/pgx"
	"go_auth_server/database"
	"go_auth_server/database/models"
	"log"
	"time"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository() *UserRepository {
	conn, err := database.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	return &UserRepository{conn}
}

func (userRepository *UserRepository) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	err := userRepository.conn.QueryRow("SELECT id, password FROM users WHERE email=$1", email).Scan(&user.ID, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) CreateUser(email, username, password string) (string, error) {
	t := time.Now()
	r, err := userRepository.conn.Exec("INSERT INTO users (email, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", email, username, password, t, t)

	return string(r), err
}