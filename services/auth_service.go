package services

import (
	"context"
	"errors"
	"go_auth_server/repositories"
	"go_auth_server/utils"

	"github.com/jackc/pgx"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{repositories.NewUserRepository()}
}

// TODO check hash password
func (as *AuthService) Auth(ctx context.Context, email string, password string) (bool, error) {
	user, err := as.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	res := utils.ComparePasswords(password, user.Password)
	return res, nil
}

// TODO hash password
func (as *AuthService) Register(ctx context.Context, email string, username string, password string) (string, error) {
	hash_password, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}
	res, err := as.userRepository.CreateUser(ctx, email, username, hash_password)
	if res != "" {
		return "OK", nil
	}
	return "", err
}
