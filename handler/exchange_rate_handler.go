package handler

import (
	"encoding/json"
	"net/http"
	"test_golang/internal/db"
)

type ExchangeRateProvider interface {
	GetExchangeRate() (float64, error)
}

type ExchangeRateAPIProvider struct{}

func (e *ExchangeRateAPIProvider) GetExchangeRate() (float64, error) {
	resp, err := http.Get("https://api.exchangerate-api.com/v4/latest/USD")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}
	rates := data["rates"].(map[string]interface{})
	exchangeRate, ok := rates["UAH"].(float64)
	if !ok {
		return 0, err
	}

	return exchangeRate, nil
}

func GetCurrentExchangeRate(database db.DatabaseInterface, exchangeRateProvider ExchangeRateProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exchangeRate, err := exchangeRateProvider.GetExchangeRate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		go func() {
			err := database.InsertExchangeRate(exchangeRate)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(exchangeRate)
	}
}
