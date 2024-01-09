package services

import (
	"context"
	"errors"
	"go_auth_server/repositories"

	"github.com/jackc/pgx"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{repositories.NewUserRepository()}
}

func (as *AuthService) Auth(ctx context.Context, email string, password string) (bool, error) {
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

func (as *AuthService) Register(ctx context.Context, email string, username string, password string) (string, error) {
	res, err := as.userRepository.CreateUser(email, username, password)
	if res != "" {
		return "OK", nil
	}
	return "", err
}
