package loginrepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
)

type LoginRepository interface {
	Validate(ctx context.Context, db *sql.DB, login *domain.Login) error
	CreateLogin(ctx context.Context, db *sql.DB, login *domain.Login) error
}
