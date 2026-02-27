package ports

import "context"

type UserChecker interface {
	Exists(ctx context.Context, userID string) (bool, error)
}
