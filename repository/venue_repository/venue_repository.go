package venuerepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
)

type VenueRepository interface {
	GetAllVenue(ctx context.Context, db *sql.DB) ([]domain.Venue, error)
	FindVenueById(ctx context.Context, db *sql.DB, idVenue string) (domain.Venue, error)
}
