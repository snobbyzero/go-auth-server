package controllers

import (
	"encoding/json"
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
	user := struct {
		Email *string `json:"email"`
		Password *string `json:"password"`
	}{}

	body := r.Body
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	resp, err := authController.userService.Auth(*user.Email, *user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(resp)
}

func (authController *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Register")
	if err != nil {
		log.Println(err)
	}
}