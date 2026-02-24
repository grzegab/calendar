package domain

import "context"

type SchedulingRepository interface {
	Save(ctx context.Context, schedule *Schedule) error
	GetById(ctx context.Context, id string) (*Schedule, error)
	List(ctx context.Context) ([]Schedule, error)
}
