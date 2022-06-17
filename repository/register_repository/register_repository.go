package registerrepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
)

type RegisterRepository interface {
	Register(ctx context.Context, db *sql.DB, register *domain.Register) error
}
