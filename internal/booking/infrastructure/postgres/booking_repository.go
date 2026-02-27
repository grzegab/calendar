package postgres

import (
	"context"
	"database/sql"
	"github/grzegab/calendar/internal/booking/domain"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Save(ctx context.Context, booking *domain.Booking) error {
	return nil
}
