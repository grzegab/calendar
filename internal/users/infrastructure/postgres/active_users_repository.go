package postgres

import (
	"database/sql"
	"github/grzegab/calendar/internal/users/application/active_user_list"
)

type ActiveUserRepository struct {
	db *sql.DB
}

func NewActiveUserRepository(db *sql.DB) *ActiveUserRepository {
	return &ActiveUserRepository{db: db}
}

func FindActive() ([]active_user_list.View, error) {
	return nil, nil
}
