package db_test

import (
	"context"
	"database/sql"
	"testing"
	"github.com/stretchr/testify/assert"
	_ "github.com/lib/pq"
	"myproject/backend/adapters/db"
	"myproject/backend/core"
)

func TestSaveSubscription(t *testing.T) {
	connStr := "user=user password=password dbname=subscriptions sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)
	defer dbConn.Close()

	repo := db.NewSubscriptionRepository(dbConn)
	sub := &core.Subscription{
		ID:    "1",
		Email: "test@example.com",
		Plan:  "premium",
	}

	err = repo.Save(context.Background(), sub)
	assert.NoError(t, err)
}
