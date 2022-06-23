package reservationrepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
)

type ReservationRepository interface {
	GetReservationSchedule(ctx context.Context, db *sql.DB, reservation *domain.Reservation) ([]int, error)
	GetUserReservationById(ctx context.Context, db *sql.DB, userId string) ([]domain.Reservation, error)
	CreateReservation(ctx context.Context, db *sql.DB, reservation *domain.Reservation) (domain.Reservation, error)
	UpdateReservation(ctx context.Context, db *sql.DB, reservation *domain.Reservation) (domain.Reservation, error)
	CancelReservation(ctx context.Context, db *sql.DB, reservationId string) error
	GetReservationScheduleForUpdate(ctx context.Context, db *sql.DB, reservation *domain.Reservation) ([]int, error)
}
