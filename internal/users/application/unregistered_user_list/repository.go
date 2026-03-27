package unregistered_user_list

type ReadRepository interface {
	FindActive() ([]View, error)
}
