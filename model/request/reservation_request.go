package request

type ReservationRequest struct {
	IdTransaction string `json:"idTransaction,omitempty"`
	UserId        string `json:"userId,omitempty"`
	VenueId       string `json:"venueId,omitempty"`
	Date          int64  `json:"date,omitempty"`
	BeginTime     int64  `json:"beginTime,omitempty"`
	EndTime       int64  `json:"endTime,omitempty"`
	Hours         []int  `json:"hours,omitempty"`
	BookingTime   int64  `json:"bookingTime,omitempty"`
	TotalPrice    int64  `json:"totalPrice,omitempty"`
}
