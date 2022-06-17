package venuecontroller

import (
	"github.com/labstack/echo"
)

type VenueController interface {
	GetAllVenue(ctx echo.Context) error
	FindVenueById(ctx echo.Context) error
}
