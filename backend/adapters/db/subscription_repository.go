package db

import (
	"database/sql"
	"context"
	_ "github.com/lib/pq"
	"myproject/backend/core"
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

func (r *SubscriptionRepository) GetByID(ctx context.Context, id string) (*core.Subscription, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, plan FROM subscriptions WHERE id = $1", id)
	sub := &core.Subscription{}
	if err := row.Scan(&sub.ID, &sub.Email, &sub.Plan); err != nil {
		return nil, err
	}
	return sub, nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, sub *core.Subscription) error {
	_, err := r.db.ExecContext(ctx, "UPDATE subscriptions SET email = $1, plan = $2 WHERE id = $3", sub.Email, sub.Plan, sub.ID)
	return err
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM subscriptions WHERE id = $1", id)
	return err
}
