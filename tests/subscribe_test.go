package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"test_golang/handler"
	_ "test_golang/handler"
	"testing"
)

type mockSubscriptionService struct{}

func (m *mockSubscriptionService) CheckEmailExists(email string) (bool, error) {
	return false, nil
}

func (m *mockSubscriptionService) InsertSubscribedUser(email string) error {
	return nil
}

func TestSubscribeHandler(t *testing.T) {
	mockService := &mockSubscriptionService{}

	reqBody, _ := json.Marshal(map[string]string{
		"email": "test@example.com",
	})
	req, err := http.NewRequest("POST", "/subscribe", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := handler.SubscribeHandler(mockService)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"message":"E-mail додано"}`
	actual := rr.Body.String()

	expected = strings.TrimSpace(expected)
	actual = strings.TrimSpace(actual)

	if actual != expected {
		t.Errorf("handler returned unexpected body:\nactual:   %v\nexpected: %v",
			actual, expected)
	}

	// Output actual and expected bodies for detailed comparison
	fmt.Println("Actual Body:", actual)
	fmt.Println("Expected Body:", expected)
}
