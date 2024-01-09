package main

import (
	"context"
	"go_auth_server/controllers"
	"go_auth_server/database"
	"go_auth_server/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Try to connect to DB")
	conn, err := database.ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to DB")
	log.Println("Ping DB")

	if err := conn.Ping(context.Background()); err != nil {
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

	log.Println("Server is starting...")
	// TODO Option pattern for http.Server
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalln(err)
	}
}
