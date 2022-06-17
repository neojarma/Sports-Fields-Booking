package domain

type Transaction struct {
	IdTransaction string
	IdVenue       string
	IdUser        string
	BeginTime     int
	EndTime       int
	hours         []int
	status        string
}
