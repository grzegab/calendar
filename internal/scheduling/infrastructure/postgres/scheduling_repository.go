package postgres

import (
	"context"
	"database/sql"
	"github/grzegab/calendar/internal/scheduling/domain"
)

type SchedulingRepository struct {
	db *sql.DB
}

func NewSchedulingRepository(db *sql.DB) *SchedulingRepository {
	return &SchedulingRepository{db: db}
}

func (r *SchedulingRepository) Save(ctx context.Context, schedule *domain.Schedule) error {
	return nil
}
