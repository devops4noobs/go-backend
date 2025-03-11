// create-users/handlers/users.go
package handlers

import (
	"backend-users/common/utils" // Import utilities
	"backend-users/create-users/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("🔥 CreateUserHandler received a request!") // Debug log

		// Check request headers
		for key, values := range r.Header {
			log.Printf("Header: %s = %s", key, values)
		}
		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			log.Println("❌ Invalid JSON:", err)
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		err := db.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", u.Name).Scan(&u.ID)
		if err != nil {
			log.Println("❌ Database Insert Error:", err)
			utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to insert user")
			return
		}

		utils.JSONResponse(w, http.StatusCreated, u)
	}
}
