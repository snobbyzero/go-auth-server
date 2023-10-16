package services

import "go_auth_server/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repositories.NewUserRepository()}
}


func (us *UserService) auth() {

}

func (us *UserService) register() {
}
