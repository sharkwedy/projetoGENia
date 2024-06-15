package application

import (
	"context"
	"myproject/backend/core"
)

type SubscriptionService struct {
	repo core.SubscriptionRepository
}

func NewSubscriptionService(repo core.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) CreateSubscription(ctx context.Context, req *CreateSubscriptionRequest) error {
	sub := &core.Subscription{ID: req.ID, Email: req.Email, Plan: req.Plan}
	return s.repo.Save(ctx, sub)
}

type CreateSubscriptionRequest struct {
	ID    string
	Email string
	Plan  string
}