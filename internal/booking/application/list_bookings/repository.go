package list_bookings

type ReadRepository interface {
	UserBookings(id string) ([]View, error)
}
