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

func (m *MockService) GetSubscription(ctx context.Context, id string) (*core.Subscription, error) {
	return &core.Subscription{ID: id, Email: "test@example.com", Plan: "premium"}, nil
}

func (m *MockService) UpdateSubscription(ctx context.Context, req *application.UpdateSubscriptionRequest) error {
	return nil
}

func (m *MockService) DeleteSubscription(ctx context.Context, id string) error {
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

func TestGetSubscriptionHandler(t *testing.T) {
	service := &MockService{}
	handler := http.NewSubscriptionHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/subscriptions/{id}", handler.GetSubscription).Methods("GET")

	req, _ := http.NewRequest("GET", "/api/subscriptions/1", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var sub core.Subscription
	err := json.NewDecoder(rr.Body).Decode(&sub)
	assert.NoError(t, err)
	assert.Equal(t, "1", sub.ID)
	assert.Equal(t, "test@example.com", sub.Email)
	assert.Equal(t, "premium", sub.Plan)
}

func TestUpdateSubscriptionHandler(t *testing.T) {
	service := &MockService{}
	handler := http.NewSubscriptionHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/subscriptions/{id}", handler.UpdateSubscription).Methods("PUT")

	body := map[string]string{
		"Email": "updated@example.com",
		"Plan":  "basic",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("PUT", "/api/subscriptions/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteSubscriptionHandler(t *testing.T) {
	service := &MockService{}
	handler := http.NewSubscriptionHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/subscriptions/{id}", handler.DeleteSubscription).Methods("DELETE")

	req, _ := http.NewRequest("DELETE", "/api/subscriptions/1", nil)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
