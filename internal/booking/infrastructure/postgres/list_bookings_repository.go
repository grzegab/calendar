package postgres

import (
	"database/sql"
	"github/grzegab/calendar/internal/booking/application/list_bookings"
)

type ListBookingsRepository struct {
	db *sql.DB
}

func NewListBookingsRepository(db *sql.DB) *ListBookingsRepository {
	return &ListBookingsRepository{db: db}
}

func (r *ListBookingsRepository) UserBookings(id string) ([]list_bookings.View, error) {
	return nil, nil
}
