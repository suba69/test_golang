package handler

import (
	"encoding/json"
	"net/http"
)

type SubscriptionService interface {
	CheckEmailExists(email string) (bool, error)
	InsertSubscribedUser(email string) error
}

func SubscribeHandler(subscriptionService SubscriptionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")

		exists, err := subscriptionService.CheckEmailExists(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}

		err = subscriptionService.InsertSubscribedUser(email)
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
