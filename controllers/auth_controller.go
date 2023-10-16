package controllers

import (
	"fmt"
	"go_auth_server/services"
	"log"
	"net/http"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{services.NewUserService()}
}

func (authController *AuthController) AuthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		log.Fatalln(err)
	}
}

func (authController *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Register")
	if err != nil {
		log.Fatalln(err)
	}
}