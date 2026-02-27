package domain

import "context"

type BookingRepository interface {
	Save(ctx context.Context, booking *Booking) error
}
