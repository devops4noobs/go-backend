package main

import (
	"backend-users/common/config"
	"backend-users/common/middleware" // Import middleware
	"backend-users/create-users/db"
	"backend-users/create-users/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware) // Apply logging middleware

	r.HandleFunc("/users", handlers.CreateUserHandler(dbConn)).Methods("POST")

	log.Println("Create Users Service running on port 8080")
	http.ListenAndServe(":8080", r)
}
