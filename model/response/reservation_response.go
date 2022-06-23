package response

type ReservationResponse struct {
	IdTransaction string `json:"idTransaction,omitempty"`
	IdVenue       string `json:"idVenue,omitempty"`
	BeginTime     int64  `json:"beginTime,omitempty"`
	EndTime       int64  `json:"endTime,omitempty"`
	Status        string `json:"status,omitempty"`
	Hours         []int  `json:"hours,omitempty"`
	BookingTime   int64  `json:"bookingTime,omitempty"`
	TotalPrice    int64  `json:"totalPrice,omitempty"`
}
