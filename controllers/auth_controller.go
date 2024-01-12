package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_auth_server/services"
	"go_auth_server/utils"
	"go_auth_server/utils/validator"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(ctx context.Context) *AuthController {
	return &AuthController{services.NewAuthService(ctx)}
}

// TODO create context with timeout
func (authController *AuthController) AuthHandler(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Email    *string `json:"email" validate:"string,required,min=5,max=100"`
		Password *string `json:"password" validate:"string,required,min=7,max=50"`
	}{}

	// Parse json body
	if err := utils.Unmarshal(&user, r.Body); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if errs := validator.Validate(user); len(errs) > 0 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Millisecond)
	defer cancel()

	res, err := authController.authService.Auth(ctx, *user.Email, *user.Password)
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

// TODO create context with timeout
func (authController *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Email    *string `json:"email" validate:"string,required,min=5,max=100"`
		Username *string `json:"username" validate:"string,required,min=5,max=50"`
		Password *string `json:"password" validate:"string,required,min=7,max=50"`
	}{}

	// Parse json body
	if err := utils.Unmarshal(&user, r.Body); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if errs := validator.Validate(user); len(errs) > 0 {
		http.Error(w, errors.Join(errs...).Error(), http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	res, err := authController.authService.Register(ctx, *user.Email, *user.Username, *user.Password)
	if err != nil {
		var pgErr *pgconn.PgError
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
