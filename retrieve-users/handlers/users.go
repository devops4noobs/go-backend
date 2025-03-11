// retrieve-users/handlers/users.go
package handlers

import (
	"database/sql"
	"net/http"
	"backend-users/retrieve-users/models"
	"backend-users/common/utils"  // Import utilities
)

func GetUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve users")
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var u models.User
			if err := rows.Scan(&u.ID, &u.Name); err != nil {
				utils.ErrorResponse(w, http.StatusInternalServerError, "Error scanning user data")
				return
			}
			users = append(users, u)
		}

		utils.JSONResponse(w, http.StatusOK, users)
	}
}
