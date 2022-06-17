package userrepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, db *sql.DB, username string) (domain.User, error)
	UpdateUser(ctx context.Context, db *sql.DB, req *domain.User) error
}
