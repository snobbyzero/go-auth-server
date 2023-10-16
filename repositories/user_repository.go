package repositories

import (
	"database/sql"
	"go_auth_server/database"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	db, err := database.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	return &UserRepository{db}
}

func (userRepository *UserRepository) getUserByEmailAndPassword() {

}