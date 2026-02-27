package user_details

type ReadRepository interface {
	GetUserDetails(id string) (View, error)
}
