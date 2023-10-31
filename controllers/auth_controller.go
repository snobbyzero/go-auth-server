package controllers

import (
	"encoding/json"
	"errors"
	"go_auth_server/services"
	"go_auth_server/utils"
	"log"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{services.NewAuthService()}
}

func (authController *AuthController) AuthHandler(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Email *string `json:"email"`
		Password *string `json:"password"`
	}{}

	// Parse json body
	if err := utils.Unmarshal(&user, r.Body); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	/*decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
	}*/

	res, err := authController.authService.Auth(*user.Email, *user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if res == false {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	// TODO return JWT tokens
	json.NewEncoder(w).Encode(res)
	return
}

// RegisterHandler TODO
func (authController *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Email *string `json:"email" validate:"string,required,min=5,max=100"`
		Username *string `json:"username" validate:"string,required,min=5,max=100"`
		Password *string `json:"password" validate:"string,required,min=5,max=100"`
	}{}

	// Parse json body
	if err := utils.Unmarshal(&user, r.Body); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	/*decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
	}*/
	if errs := utils.Validate(user); len(errs) > 0 {
		http.Error(w, errors.Join(errs...).Error(), http.StatusInternalServerError)
		return
	}

	res, err := authController.authService.Register(*user.Email, *user.Username, *user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
	return
}