package reservationcontroller

import (
	"github.com/labstack/echo"
)

type ReservationController interface {
	GetReservationSchedule(ctx echo.Context) error
	GetUserReservationById(ctx echo.Context) error
	CreateReservation(ctx echo.Context) error
	UpdateReservation(ctx echo.Context) error
	CancelReservation(ctx echo.Context) error
	GetReservationScheduleForUpdate(ctx echo.Context) error
}
