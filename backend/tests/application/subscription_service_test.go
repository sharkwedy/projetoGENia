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
