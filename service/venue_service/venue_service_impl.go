package venueservice

import (
	"booking_fields/model/response"
	venuerepository "booking_fields/repository/venue_repository"
	"context"
	"database/sql"
)

type VenueServiceImpl struct {
	VenuRepository venuerepository.VenueRepository
	Db             *sql.DB
}

func NewVenueService(repo venuerepository.VenueRepository, db *sql.DB) VenueService {
	return &VenueServiceImpl{
		VenuRepository: repo,
		Db:             db,
	}
}

func (service *VenueServiceImpl) GetAllVenue(ctx context.Context) ([]response.VenueResponse, error) {
	domainResult, err := service.VenuRepository.GetAllVenue(ctx, service.Db)

	if err != nil {
		return nil, err
	}

	result := []response.VenueResponse{}

	for _, each := range domainResult {
		res := response.VenueResponse{
			IdVenue:   each.IdVenue,
			Title:     each.Title,
			Location:  each.Location,
			Price:     each.Price,
			Category:  each.Category,
			Rating:    each.Rating,
			ImageLink: each.ImageLink,
		}

		result = append(result, res)
	}

	return result, nil

}

func (service *VenueServiceImpl) FindVenueById(ctx context.Context, venueId string) (response.VenueResponse, error) {
	domainResult, err := service.VenuRepository.FindVenueById(ctx, service.Db, venueId)

	if err != nil {
		return response.VenueResponse{}, err
	}

	return response.VenueResponse{
		IdVenue:   domainResult.IdVenue,
		Title:     domainResult.Title,
		Location:  domainResult.Location,
		Price:     domainResult.Price,
		Category:  domainResult.Category,
		Rating:    domainResult.Rating,
		ImageLink: domainResult.ImageLink,
	}, nil
}
