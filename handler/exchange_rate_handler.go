package handler

import (
	"encoding/json"
	"net/http"
	"test_golang/db"
)

func GetExchangeRateFromAPI() (float64, error) {
	resp, err := http.Get("https://open.er-api.com/v6/latest/USD")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}
	exchangeRate := data["rates"].(map[string]interface{})["UAH"].(float64)

	return exchangeRate, nil
}

func GetCurrentExchangeRate(database *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exchangeRate, err := GetExchangeRateFromAPI()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = database.InsertExchangeRate(exchangeRate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":       "Exchange rate updated successfully",
			"exchange_rate": exchangeRate,
		})
	}
}
