package list_bookings

type Repository interface {
	UserBookings(id string) ([]View, error)
}
