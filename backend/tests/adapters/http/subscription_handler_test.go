package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"myproject/backend/adapters/http"
	"myproject/backend/application"
)

type MockService struct {}

func (m *MockService) CreateSubscription(ctx context.Context, req *application.CreateSubscriptionRequest) error {
	return nil
}

func TestCreateSubscriptionHandler(t *testing.T) {
	service := &MockService{}
	handler := http.NewSubscriptionHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/subscriptions", handler.CreateSubscription).Methods("POST")

	body := map[string]string{
		"ID":    "1",
		"Email": "test@example.com",
		"Plan":  "premium",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/subscriptions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
