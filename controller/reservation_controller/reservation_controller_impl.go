package reservationcontroller

import (
	"booking_fields/model/request"
	"booking_fields/model/response"
	reservationservice "booking_fields/service/reservation_service"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ReservationControllerImpl struct {
	ReservationService reservationservice.ReservationService
}

func NewReservationController(service reservationservice.ReservationService) ReservationController {
	return &ReservationControllerImpl{
		ReservationService: service,
	}
}

func (controller *ReservationControllerImpl) GetReservationSchedule(ctx echo.Context) error {
	type Response struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data   *[]int `json:"data,omitempty"`
	}
	ctxBack := context.Background()

	venueId := ctx.QueryParam("venue")
	dateParam := ctx.QueryParam("date")
	date, err := strconv.Atoi(dateParam)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
	}

	req := &request.ReservationRequest{
		VenueId: venueId,
		Date:    int64(date),
	}

	result, err := controller.ReservationService.GetReservationSchedule(ctxBack, req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "succes get data schedule",
		Data:   &result,
	})
}

func (controller *ReservationControllerImpl) GetUserReservationById(ctx echo.Context) error {
	type Response struct {
		Code   int                             `json:"code"`
		Status string                          `json:"status"`
		Data   *[]response.ReservationResponse `json:"data,omitempty"`
	}

	reserveId := ctx.Param("id")
	ctxBack := context.Background()

	result, err := controller.ReservationService.GetUserReservationById(ctxBack, reserveId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "succes get data schedule",
		Data:   &result,
	})
}

func (controller *ReservationControllerImpl) CreateReservation(ctx echo.Context) error {
	type Response struct {
		Code   int                           `json:"code"`
		Status string                        `json:"status"`
		Data   *response.ReservationResponse `json:"data,omitempty"`
	}

	ctxBack := context.Background()
	req := new(request.ReservationRequest)
	ctx.Bind(req)

	result, err := controller.ReservationService.CreateReservation(ctxBack, req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   &result,
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "success create reservation",
		Data:   &result,
	})
}

func (controller *ReservationControllerImpl) UpdateReservation(ctx echo.Context) error {
	type Response struct {
		Code   int                           `json:"code"`
		Status string                        `json:"status"`
		Data   *response.ReservationResponse `json:"data,omitempty"`
	}

	ctxBack := context.Background()
	req := new(request.ReservationRequest)
	ctx.Bind(req)

	result, err := controller.ReservationService.UpdateReservation(ctxBack, req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   &result,
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "success update reservation",
		Data:   &result,
	})
}

func (controller *ReservationControllerImpl) CancelReservation(ctx echo.Context) error {
	type Response struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
	}

	ctxBack := context.Background()
	reservationId := ctx.Param("id")

	err := controller.ReservationService.CancelReservation(ctxBack, reservationId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "success delete reservation",
	})
}

func (controller *ReservationControllerImpl) GetReservationScheduleForUpdate(ctx echo.Context) error {
	type Response struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data   *[]int `json:"data,omitempty"`
	}
	ctxBack := context.Background()

	venueId := ctx.QueryParam("venue")
	dateParam := ctx.QueryParam("date")
	idTransaction := ctx.QueryParam("txId")
	date, err := strconv.Atoi(dateParam)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
	}

	req := &request.ReservationRequest{
		VenueId:       venueId,
		Date:          int64(date),
		IdTransaction: idTransaction,
	}

	result, err := controller.ReservationService.GetReservationScheduleForUpdate(ctxBack, req)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "succes get data schedule",
		Data:   &result,
	})
}
