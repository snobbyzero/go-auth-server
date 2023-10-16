package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go_auth_server/controllers"
	"go_auth_server/database"
	"log"
	"net/http"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	authController := controllers.NewAuthController()
	router := mux.NewRouter()
	//router.StrictSlash(true)
	router.HandleFunc("/auth", authController.AuthHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", authController.RegisterHandler).Methods(http.MethodPost)

	err = http.ListenAndServe(":8081", router)

	if err != nil {
		panic(err)
	}
}
