package infrastructure

import (
	"context"
	"database/sql"
	"errors"

	"github/grzegab/calendar/internal/users/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	// SQL insert
	return nil
}

func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (r *UserRepository) FindByPhone(phone string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (r *UserRepository) findOne(
	ctx context.Context,
	where string,
	arg any,
) (*domain.User, error) {

	query := `
        SELECT id, email, phone, status
        FROM users
        WHERE ` + where + `
        LIMIT 1
    `

	row := r.db.QueryRowContext(ctx, query, arg)

	var (
		id     string
		email  string
		phone  string
		status string
	)

	if err := row.Scan(&id, &email, &phone, &status); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrorUserNotFound
		}
		return nil, err
	}

	return mapToDomain(id, email, phone, status)
}

func mapToDomain(
	id string,
	email string,
	phone string,
	status string,
) (*domain.User, error) {
	e, err := domain.NewEmail(email)
	if err != nil {
		return nil, err
	}

	p, err := domain.NewPhoneNumber(phone)
	if err != nil {
		return nil, err
	}

	user := domain.RehydrateUser(
		domain.UserID(id),
		e,
		p,
		domain.UserStatusFromString(status),
	)

	return user, nil
}
