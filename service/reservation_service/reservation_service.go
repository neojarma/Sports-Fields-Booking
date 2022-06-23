package reservationservice

import (
	"booking_fields/model/request"
	"booking_fields/model/response"
	"context"
)

type ReservationService interface {
	GetReservationSchedule(ctx context.Context, request *request.ReservationRequest) ([]int, error)
	GetUserReservationById(ctx context.Context, userId string) ([]response.ReservationResponse, error)
	CreateReservation(ctx context.Context, request *request.ReservationRequest) (response.ReservationResponse, error)
	UpdateReservation(ctx context.Context, request *request.ReservationRequest) (response.ReservationResponse, error)
	CancelReservation(ctx context.Context, reservationId string) error
	GetReservationScheduleForUpdate(ctx context.Context, request *request.ReservationRequest) ([]int, error)
}
