package adapters

import (
	"context"
	"errors"
	"github/grzegab/calendar/internal/users/domain"
)

type BookingUserChecker struct {
	repo domain.Repository
}

func NewBookingUserChecker(repo domain.Repository) *BookingUserChecker {
	return &BookingUserChecker{repo: repo}
}

func (a *BookingUserChecker) Exists(
	ctx context.Context,
	userID string,
) (bool, error) {

	_, err := a.repo.GetByID(userID)
	if err != nil {
		if errors.Is(err, domain.ErrorUserNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
