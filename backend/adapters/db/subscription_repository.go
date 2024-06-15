package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"context"
)

type SubscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Save(ctx context.Context, sub *core.Subscription) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO subscriptions (id, email, plan) VALUES ($1, $2, $3)", sub.ID, sub.Email, sub.Plan)
	return err
}