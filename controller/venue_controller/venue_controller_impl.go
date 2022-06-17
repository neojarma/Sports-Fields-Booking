package venuecontroller

import (
	"booking_fields/model/response"
	venueservice "booking_fields/service/venue_service"
	"context"
	"net/http"

	"github.com/labstack/echo"
)

type VenueControllerImpl struct {
	VenueService venueservice.VenueService
}

func NewVenueController(service venueservice.VenueService) VenueController {
	return &VenueControllerImpl{
		VenueService: service,
	}
}

func (controller *VenueControllerImpl) GetAllVenue(ctx echo.Context) error {
	type Response struct {
		Code   int                      `json:"code"`
		Status string                   `json:"status"`
		Data   []response.VenueResponse `json:"data"`
	}
	ctxBack := context.Background()
	result, err := controller.VenueService.GetAllVenue(ctxBack)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "succes get all data",
		Data:   result,
	})
}

func (controller *VenueControllerImpl) FindVenueById(ctx echo.Context) error {
	type Response struct {
		Code   int                     `json:"code"`
		Status string                  `json:"status"`
		Data   *response.VenueResponse `json:"data,omitempty"`
	}
	ctxBack := context.Background()

	venueId := ctx.Param("id")

	result, err := controller.VenueService.FindVenueById(ctxBack, venueId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, &Response{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "succes get all data",
		Data:   &result,
	})
}
