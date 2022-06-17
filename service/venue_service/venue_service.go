package venueservice

import (
	"booking_fields/model/response"
	"context"
)

type VenueService interface {
	GetAllVenue(ctx context.Context) ([]response.VenueResponse, error)
	FindVenueById(ctx context.Context, venueId string) (response.VenueResponse, error)
}
