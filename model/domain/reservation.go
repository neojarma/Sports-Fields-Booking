package domain

type Reservation struct {
	IdTransaction string
	UserId        string
	VenueId       string
	Date          int64
	BeginTime     int64
	EndTime       int64
	Status        string
	Hours         []int
	BookingTime   int64
	TotalPrice    int64
}
