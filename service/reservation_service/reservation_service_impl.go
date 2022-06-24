package reservationservice

import (
	"booking_fields/helper"
	"booking_fields/model/domain"
	"booking_fields/model/request"
	"booking_fields/model/response"
	reservationrepository "booking_fields/repository/reservation_repository"
	"context"
	"database/sql"
	"strconv"
)

const END_HOUR_IN_MILLIS = 39600000
const START_HOUR_IN_MILLIS = 28800000

type ReservationServiceImpl struct {
	ReservationRepository reservationrepository.ReservationRepository
	Db                    *sql.DB
}

func NewReservationService(repo reservationrepository.ReservationRepository, db *sql.DB) ReservationService {
	return &ReservationServiceImpl{
		ReservationRepository: repo,
		Db:                    db,
	}
}

func (service *ReservationServiceImpl) GetReservationSchedule(ctx context.Context, request *request.ReservationRequest) ([]int, error) {

	// front-end request just date, like YYYY_MM DD 00.00.00
	// add +8 hour
	MinDate := request.Date + START_HOUR_IN_MILLIS
	// end hour is 7 PM so we add +11 hour from star hour
	MaxDate := request.Date + START_HOUR_IN_MILLIS + END_HOUR_IN_MILLIS

	req := domain.Reservation{
		VenueId:   request.VenueId,
		BeginTime: MinDate,
		EndTime:   MaxDate,
	}

	result, err := service.ReservationRepository.GetReservationSchedule(ctx, service.Db, &req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *ReservationServiceImpl) GetUserReservationById(ctx context.Context, userId string) ([]response.ReservationResponse, error) {

	domainResult, err := service.ReservationRepository.GetUserReservationById(ctx, service.Db, userId)

	if err != nil {
		return nil, err
	}

	result := []response.ReservationResponse{}

	for _, v := range domainResult {
		item := response.ReservationResponse{
			IdTransaction: v.IdTransaction,
			IdVenue:       v.VenueId,
			BeginTime:     v.BeginTime,
			EndTime:       v.EndTime,
			Status:        v.Status,
			Hours:         v.Hours,
			BookingTime:   v.BookingTime,
			TotalPrice:    v.TotalPrice,
		}

		result = append(result, item)
	}

	return result, nil

}

func (service *ReservationServiceImpl) CreateReservation(ctx context.Context, request *request.ReservationRequest) (response.ReservationResponse, error) {

	idTransaction := "TX-" + strconv.Itoa(helper.GenerateRandomId())

	reservation := domain.Reservation{
		IdTransaction: idTransaction,
		UserId:        request.UserId,
		VenueId:       request.VenueId,
		Date:          request.Date,
		BeginTime:     request.BeginTime,
		Hours:         request.Hours,
		EndTime:       request.EndTime,
		BookingTime:   request.BookingTime,
		TotalPrice:    request.TotalPrice,
	}

	domainResult, err := service.ReservationRepository.CreateReservation(ctx, service.Db, &reservation)
	if err != nil {
		return response.ReservationResponse{}, err
	}

	return response.ReservationResponse{
		IdTransaction: domainResult.IdTransaction,
		IdVenue:       domainResult.VenueId,
		BeginTime:     domainResult.BeginTime,
		EndTime:       domainResult.EndTime,
		Hours:         domainResult.Hours,
		BookingTime:   domainResult.BookingTime,
		TotalPrice:    domainResult.TotalPrice,
	}, nil
}

func (service *ReservationServiceImpl) UpdateReservation(ctx context.Context, request *request.ReservationRequest) (response.ReservationResponse, error) {

	reservation := domain.Reservation{
		IdTransaction: request.IdTransaction,
		BeginTime:     request.BeginTime,
		Hours:         request.Hours,
		EndTime:       request.EndTime,
		BookingTime:   request.BookingTime,
		TotalPrice:    request.TotalPrice,
	}

	domainResult, err := service.ReservationRepository.UpdateReservation(ctx, service.Db, &reservation)
	if err != nil {
		return response.ReservationResponse{}, err
	}

	return response.ReservationResponse{
		IdTransaction: domainResult.IdTransaction,
		BeginTime:     domainResult.BeginTime,
		EndTime:       domainResult.EndTime,
		Hours:         domainResult.Hours,
		BookingTime:   domainResult.BookingTime,
		TotalPrice:    domainResult.TotalPrice,
	}, nil
}

func (service *ReservationServiceImpl) CancelReservation(ctx context.Context, reservationId string) error {

	err := service.ReservationRepository.CancelReservation(ctx, service.Db, reservationId)

	if err != nil {
		return err
	}

	return nil
}

func (service *ReservationServiceImpl) GetReservationScheduleForUpdate(ctx context.Context, request *request.ReservationRequest) ([]int, error) {
	// front-end request just date, like YYYY_MM DD 00.00.00
	// add +8 hour
	MinDate := request.Date + START_HOUR_IN_MILLIS
	// end hour is 7 PM so we add +11 hour from star hour
	MaxDate := request.Date + START_HOUR_IN_MILLIS + END_HOUR_IN_MILLIS

	req := domain.Reservation{
		VenueId:       request.VenueId,
		BeginTime:     MinDate,
		EndTime:       MaxDate,
		IdTransaction: request.IdTransaction,
	}

	result, err := service.ReservationRepository.GetReservationScheduleForUpdate(ctx, service.Db, &req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
