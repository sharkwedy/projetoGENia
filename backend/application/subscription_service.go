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

func (s *SubscriptionService) GetSubscription(ctx context.Context, id string) (*core.Subscription, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *SubscriptionService) UpdateSubscription(ctx context.Context, req *UpdateSubscriptionRequest) error {
	sub := &core.Subscription{ID: req.ID, Email: req.Email, Plan: req.Plan}
	return s.repo.Update(ctx, sub)
}

func (s *SubscriptionService) DeleteSubscription(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// Request structs
type CreateSubscriptionRequest struct {
	ID    string
	Email string
	Plan  string
}

type UpdateSubscriptionRequest struct {
	ID    string
	Email string
	Plan  string
}
