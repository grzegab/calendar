package postgres

import "database/sql"

type UserDetailsRepository struct {
	db *sql.DB
}

func NewUserDetailsRepository(db *sql.DB) *UserDetailsRepository {
	return &UserDetailsRepository{db: db}
}
