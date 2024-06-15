package core

type Subscription struct {
	ID    string
	Email string
	Plan  string
}

type SubscriptionRepository interface {
	Save(ctx context.Context, sub *Subscription) error
	GetByID(ctx context.Context, id string) (*Subscription, error)
	Update(ctx context.Context, sub *Subscription) error
	Delete(ctx context.Context, id string) error
}