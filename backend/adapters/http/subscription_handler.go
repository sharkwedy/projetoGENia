package http

import (
	"encoding/json"
	"net/http"
	"context"
	"github.com/gorilla/mux"
	"myproject/backend/application"
)

type SubscriptionHandler struct {
	service *application.SubscriptionService
}

func NewSubscriptionHandler(service *application.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

func (h *SubscriptionHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var req application.CreateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateSubscription(context.Background(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}