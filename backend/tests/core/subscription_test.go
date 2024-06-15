package core_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"myproject/backend/core"
)

func TestSubscription(t *testing.T) {
	sub := &core.Subscription{
		ID:    "1",
		Email: "test@example.com",
		Plan:  "premium",
	}

	assert.Equal(t, "1", sub.ID)
	assert.Equal(t, "test@example.com", sub.Email)
	assert.Equal(t, "premium", sub.Plan)
}
