package main

import (
	"context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go_auth_server/controllers"
	"go_auth_server/database"
	"go_auth_server/middleware"
	"log"
	"net/http"
)

func main() {
	conn, err := database.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	if err := database.CreateTables(); err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	authController := controllers.NewAuthController()
	router := mux.NewRouter()
	//router.StrictSlash(true)
	router.HandleFunc("/auth", authController.AuthHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", authController.RegisterHandler).Methods(http.MethodPost)

	router.Use(middleware.Logging, middleware.SetHeaders)

	// TODO Option pattern for http.Server
	err = http.ListenAndServe(":8081", router)

	if err != nil {
		log.Fatalln(err)
	}
}
