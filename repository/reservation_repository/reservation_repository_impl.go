package reservationrepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
)

type ReservationRepositoryImpl struct {
}

func NewReservationRepository() ReservationRepository {
	return &ReservationRepositoryImpl{}
}

func (repository *ReservationRepositoryImpl) GetReservationSchedule(ctx context.Context, db *sql.DB, reservation *domain.Reservation) ([]int, error) {

	SQL := "SELECT array_to_json(hours::int[]) FROM public.reservation WHERE id_venue = ($1) AND status = 'valid' AND begin_time >= ($2) AND end_time <= ($3)"

	rows, err := db.QueryContext(ctx, SQL, reservation.VenueId, reservation.BeginTime, reservation.EndTime)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allHours := []int{}

	for rows.Next() {
		// array_to_json() return []byte
		var byteJson []byte
		err := rows.Scan(&byteJson)
		if err != nil {
			return nil, err
		}

		hour, err := parseHour(byteJson)
		if err != nil {
			return nil, err
		}

		for _, v := range hour {
			allHours = append(allHours, v)
		}
	}

	return allHours, nil
}

func parseHour(byteJson []byte) ([]int, error) {

	var hour []int
	err := json.Unmarshal(byteJson, &hour)
	if err != nil {
		return nil, err
	}

	return hour, nil
}

func (repository *ReservationRepositoryImpl) GetUserReservationById(ctx context.Context, db *sql.DB, userId string) ([]domain.Reservation, error) {

	SQL := "SELECT id_transaction, id_venue, begin_time, end_time, status, array_to_json(hours::int[]), booking_time, total_price from public.reservation WHERE id_user = ($1)"

	rows, err := db.QueryContext(ctx, SQL, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := []domain.Reservation{}

	for rows.Next() {
		var byteJson []byte
		var response domain.Reservation

		err := rows.Scan(&response.IdTransaction, &response.VenueId, &response.BeginTime, &response.EndTime, &response.Status, &byteJson, &response.BookingTime, &response.TotalPrice)
		if err != nil {
			return nil, err
		}

		response.Hours, err = parseHour(byteJson)
		if err != nil {
			return nil, err
		}

		result = append(result, response)
	}

	return result, nil
}

func (repository *ReservationRepositoryImpl) CreateReservation(ctx context.Context, db *sql.DB, reservation *domain.Reservation) (domain.Reservation, error) {

	SQL := "INSERT INTO public.reservation(id_transaction, id_venue, id_user, begin_time, end_time, hours, booking_time, total_price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err := db.ExecContext(ctx, SQL, reservation.IdTransaction, reservation.VenueId, reservation.UserId, reservation.BeginTime, reservation.EndTime, pq.Array(reservation.Hours), reservation.BookingTime, reservation.TotalPrice)

	if err != nil {
		return domain.Reservation{}, err
	}

	return domain.Reservation{
		IdTransaction: reservation.IdTransaction,
		VenueId:       reservation.VenueId,
		UserId:        reservation.UserId,
		BeginTime:     reservation.BeginTime,
		EndTime:       reservation.EndTime,
		Hours:         reservation.Hours,
		BookingTime:   reservation.BookingTime,
		TotalPrice:    reservation.TotalPrice,
	}, nil
}

func (repository *ReservationRepositoryImpl) UpdateReservation(ctx context.Context, db *sql.DB, reservation *domain.Reservation) (domain.Reservation, error) {

	SQL := "UPDATE public.reservation SET begin_time = ($1), end_time = ($2), hours = ($3), booking_time = ($4), total_price = ($5) WHERE id_transaction = ($6)"

	_, err := db.ExecContext(ctx, SQL, reservation.BeginTime, reservation.EndTime, pq.Array(reservation.Hours), reservation.BookingTime, reservation.TotalPrice, reservation.IdTransaction)
	if err != nil {
		return domain.Reservation{}, err
	}

	return domain.Reservation{
		IdTransaction: reservation.IdTransaction,
		VenueId:       reservation.VenueId,
		BeginTime:     reservation.BeginTime,
		EndTime:       reservation.EndTime,
		Hours:         reservation.Hours,
		BookingTime:   reservation.BookingTime,
		TotalPrice:    reservation.TotalPrice,
	}, nil
}

func (repository *ReservationRepositoryImpl) CancelReservation(ctx context.Context, db *sql.DB, reservationId string) error {

	SQL := "UPDATE public.reservation SET status = 'invalid' WHERE id_transaction = ($1)"

	_, err := db.ExecContext(ctx, SQL, reservationId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ReservationRepositoryImpl) GetReservationScheduleForUpdate(ctx context.Context, db *sql.DB, reservation *domain.Reservation) ([]int, error) {
	SQL := "SELECT array_to_json(hours::int[]) FROM public.reservation WHERE id_venue = ($1) AND status = 'valid' AND begin_time >= ($2) AND end_time <= ($3) AND id_transaction != ($4)"

	rows, err := db.QueryContext(ctx, SQL, reservation.VenueId, reservation.BeginTime, reservation.EndTime, reservation.IdTransaction)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allHours := []int{}

	for rows.Next() {
		// array_to_json() return []byte
		var byteJson []byte
		err := rows.Scan(&byteJson)
		if err != nil {
			return nil, err
		}

		hour, err := parseHour(byteJson)
		if err != nil {
			return nil, err
		}

		for _, v := range hour {
			allHours = append(allHours, v)
		}
	}

	return allHours, nil
}
