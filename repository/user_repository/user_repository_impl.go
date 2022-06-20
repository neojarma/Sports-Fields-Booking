package userrepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) GetUserByUsername(ctx context.Context, db *sql.DB, username string) (domain.User, error) {

	SQL := "SELECT id_user, name, address, phone_number, email, encode(image_profile, 'base64') FROM public.user WHERE username = ($1)"
	rows, err := db.QueryContext(ctx, SQL, username)

	response := domain.User{}

	if err != nil {
		return response, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&response.IdUser, &response.Name, &response.Address, &response.PhoneNumber, &response.Email, &response.ImageProfile)

		if err != nil {
			return response, err
		}
	} else {
		return response, errors.New("user not found")
	}

	return response, nil
}

func (repository *UserRepositoryImpl) UpdateUser(ctx context.Context, db *sql.DB, req *domain.User) error {

	SQL := "UPDATE public.user SET name = ($1), address = ($2), phone_number = ($3), email = ($4), image_profile = decode($5, 'base64') WHERE id_user = ($6) "
	_, err := db.ExecContext(ctx, SQL, req.Name, req.Address, req.PhoneNumber, req.Email, req.ImageProfile, req.IdUser)

	if err != nil {
		return err
	}

	return nil
}
