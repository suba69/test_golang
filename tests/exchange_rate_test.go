package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"test_golang/handler"
	"test_golang/internal/db/mock"
	"testing"
)

type mockExchangeRateProvider struct {
	exchangeRate float64
	err          error
}

func (m *mockExchangeRateProvider) GetExchangeRate() (float64, error) {
	return m.exchangeRate, m.err
}

func TestGetCurrentExchangeRate(t *testing.T) {
	t.Run("Successful API Call", func(t *testing.T) {
		mockExchangeRate := 28.5
		mockProvider := &mockExchangeRateProvider{
			exchangeRate: mockExchangeRate,
		}

		mockDB := &mock.MockDatabase{
			InsertExchangeRateFn: func(rate float64) error {
				return nil
			},
		}

		req := httptest.NewRequest("GET", "/current-exchange-rate", nil)
		w := httptest.NewRecorder()

		handler := handler.GetCurrentExchangeRate(mockDB, mockProvider)
		handler.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
		}

		var responseExchangeRate float64
		err := json.NewDecoder(w.Body).Decode(&responseExchangeRate)
		if err != nil {
			t.Errorf("failed to decode response body: %v", err)
		}

		if responseExchangeRate != mockExchangeRate {
			t.Errorf("handler returned unexpected exchange rate: got %v want %v", responseExchangeRate, mockExchangeRate)
		}
	})

	t.Run("Failed API Call", func(t *testing.T) {
		mockErr := errors.New("API call failed")
		mockProvider := &mockExchangeRateProvider{
			err: mockErr,
		}

		mockDB := &mock.MockDatabase{
			InsertExchangeRateFn: func(rate float64) error {
				return nil
			},
		}

		req := httptest.NewRequest("GET", "/current-exchange-rate", nil)
		w := httptest.NewRecorder()

		handler := handler.GetCurrentExchangeRate(mockDB, mockProvider)
		handler.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
		}

		expectedErrMsg := "API call failed"
		if !strings.Contains(w.Body.String(), expectedErrMsg) {
			t.Errorf("handler returned unexpected error message: got %v want %v", w.Body.String(), expectedErrMsg)
		}
	})
}
