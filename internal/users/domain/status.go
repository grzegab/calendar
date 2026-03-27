package domain

type UserStatus int

const (
	StatusPending UserStatus = iota
	StatusActive
	StatusInactive
)

func (us UserStatus) String() string {
	switch us {
	case StatusActive:
		return "active"
	case StatusInactive:
		return "inactive"
	case StatusPending:
		return "pending"
	default:
		return "unknown"
	}
}

func UserStatusFromString(s string) UserStatus {
	switch s {
	case "active":
		return StatusActive
	case "inactive":
		return StatusInactive
	case "pending":
		return StatusPending
	default:
		return StatusPending
	}
}
