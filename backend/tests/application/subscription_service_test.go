package application_test

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
	"myproject/backend/application"
	"myproject/backend/core"
)

type MockRepository struct{}

func (m *MockRepository) Save(ctx context.Context, sub *core.Subscription) error {
	return nil
}

func (m *MockRepository) GetByID(ctx context.Context, id string) (*core.Subscription, error) {
	return &core.Subscription{ID: id, Email: "test@example.com", Plan: "premium"}, nil
}

func (m *MockRepository) Update(ctx context.Context, sub *core.Subscription) error {
	return nil
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func TestCreateSubscription(t *testing.T) {
	repo := &MockRepository{}
	service := application.NewSubscriptionService(repo)

	req := &application.CreateSubscriptionRequest{
		ID:    "1",
		Email: "test@example.com",
		Plan:  "premium",
	}

	err := service.CreateSubscription(context.Background(), req)
	assert.NoError(t, err)
}

func TestGetSubscription(t *testing.T) {
	repo := &MockRepository{}
	service := application.NewSubscriptionService(repo)

	sub, err := service.GetSubscription(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", sub.ID)
	assert.Equal(t, "test@example.com", sub.Email)
	assert.Equal(t, "premium", sub.Plan)
}

func TestUpdateSubscription(t *testing.T) {
	repo := &MockRepository{}
	service := application.NewSubscriptionService(repo)

	req := &application.UpdateSubscriptionRequest{
		ID:    "1",
		Email: "updated@example.com",
		Plan:  "basic",
	}

	err := service.UpdateSubscription(context.Background(), req)
	assert.NoError(t, err)
}

func TestDeleteSubscription(t *testing.T) {
	repo := &MockRepository{}
	service := application.NewSubscriptionService(repo)

	err := service.DeleteSubscription(context.Background(), "1")
	assert.NoError(t, err)
}
