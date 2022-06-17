package response

type VenueResponse struct {
	IdVenue   string  `json:"idVenue"`
	Title     string  `json:"title"`
	Location  string  `json:"location"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
	Rating    float64 `json:"rating"`
	ImageLink string  `json:"image"`
}
