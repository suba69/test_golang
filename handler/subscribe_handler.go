package handler

import (
	"encoding/json"
	"net/http"
	"test_golang/db"
	_ "test_golang/entity"
)

func SubscribeHandler(database *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")

		exists, err := database.CheckEmailExists(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}

		err = database.InsertSubscribedUser(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Email subscribed successfully",
			"email":   email,
		})
	}
}
