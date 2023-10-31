package services

import (
	"errors"
	"github.com/jackc/pgx"
	"go_auth_server/repositories"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{repositories.NewUserRepository()}
}


func (as *AuthService) Auth(email, password string) (bool, error) {
	user, err := as.userRepository.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	if user.Password == password {
		return true, nil
	}
	return false, nil
}

func (as *AuthService) Register(email, username, password string) (string, error) {
	res, err := as.userRepository.CreateUser(email, username, password)

	return res, err
}
