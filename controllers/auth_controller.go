package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_auth_server/services"
	"go_auth_server/utils"
	"log"
	"net/http"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{services.NewAuthService()}
}

func (authController *AuthController) AuthHandler(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Email    *string `json:"email"`
		Password *string `json:"password"`
	}{}

	// Parse json body
	if err := utils.Unmarshal(&user, r.Body); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res, err := authController.authService.Auth(r.Context(), *user.Email, *user.Password)
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
		Email    *string `json:"email" validate:"string,required,min=5,max=100"`
		Username *string `json:"username" validate:"string,required,min=5,max=100"`
		Password *string `json:"password" validate:"string,required,min=5,max=100"`
	}{}

	// Parse json body
	if err := utils.Unmarshal(&user, r.Body); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if errs := utils.Validate(user); len(errs) > 0 {
		http.Error(w, errors.Join(errs...).Error(), http.StatusInternalServerError)
		return
	}

	res, err := authController.authService.Register(r.Context(), *user.Email, *user.Username, *user.Password)
	if err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			var field string
			switch pgErr.ConstraintName {
			case "users_email_key":
				field = "email"
			case "users_username_key":
				field = "username"
			default:
				field = pgErr.ConstraintName
			}
			log.Println(pgErr)
			http.Error(w, fmt.Sprintf("User with this %s already exists", field), http.StatusOK)

		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(res)
	return
}
