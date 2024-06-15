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

func TestGetSubscriptionByID(t *testing.T) {
	connStr := "user=user password=password dbname=subscriptions sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)
	defer dbConn.Close()

	repo := db.NewSubscriptionRepository(dbConn)

	sub, err := repo.GetByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", sub.ID)
	assert.Equal(t, "test@example.com", sub.Email)
	assert.Equal(t, "premium", sub.Plan)
}

func TestUpdateSubscription(t *testing.T) {
	connStr := "user=user password=password dbname=subscriptions sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)
	defer dbConn.Close()

	repo := db.NewSubscriptionRepository(dbConn)
	sub := &core.Subscription{
		ID:    "1",
		Email: "updated@example.com",
		Plan:  "basic",
	}

	err = repo.Update(context.Background(), sub)
	assert.NoError(t, err)
}

func TestDeleteSubscription(t *testing.T) {
	connStr := "user=user password=password dbname=subscriptions sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)
	defer dbConn.Close()

	repo := db.NewSubscriptionRepository(dbConn)

	err = repo.Delete(context.Background(), "1")
	assert.NoError(t, err)
}
