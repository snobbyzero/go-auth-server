package services

import (
	"go_auth_server/repositories"
	"log"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repositories.NewUserRepository()}
}


func (us *UserService) Auth(email string, password string) (bool, error) {
	user, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if user.Password == password {
		return true, nil
	}
	return false, nil
}

func (us *UserService) Register() {
}
