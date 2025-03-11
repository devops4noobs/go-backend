package main

import (
	"backend-users/common/config"
	"backend-users/common/middleware" // Import middleware
	"backend-users/retrieve-users/db"
	"backend-users/retrieve-users/handlers"
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

	r.HandleFunc("/getUsers", handlers.GetUsersHandler(dbConn)).Methods("GET")

	log.Println("Retrieve Users Service running on port 8085")
	http.ListenAndServe(":8085", r)
}
