package venuerepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
	"errors"
)

type VenueRepositoryImpl struct {
}

func NewVenueRepository() VenueRepository {
	return &VenueRepositoryImpl{}
}

func (repository *VenueRepositoryImpl) GetAllVenue(ctx context.Context, db *sql.DB) ([]domain.Venue, error) {

	SQL := "SELECT id_venue, title, location, price, category, rating, image_link FROM public.venue"
	rows, err := db.QueryContext(ctx, SQL)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := []domain.Venue{}

	for rows.Next() {
		venue := domain.Venue{}

		err := rows.Scan(&venue.IdVenue, &venue.Title, &venue.Location, &venue.Price, &venue.Category, &venue.Rating, &venue.ImageLink)

		if err != nil {
			return nil, err
		}

		result = append(result, venue)
	}

	return result, nil
}

func (repository *VenueRepositoryImpl) FindVenueById(ctx context.Context, db *sql.DB, idVenue string) (domain.Venue, error) {

	SQL := "SELECT id_venue, title, location, price, category, rating, image_link FROM public.venue WHERE id_venue = ($1)"
	rows, err := db.QueryContext(ctx, SQL, idVenue)

	venue := domain.Venue{}

	if err != nil {
		return venue, err
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&venue.IdVenue, &venue.Title, &venue.Location, &venue.Price, &venue.Category, &venue.Rating, &venue.ImageLink)

		if err != nil {
			return venue, err
		}
	} else {
		return venue, errors.New("venue not found")
	}

	return venue, nil

}
