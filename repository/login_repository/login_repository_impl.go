package loginrepository

import (
	"booking_fields/model/domain"
	"context"
	"database/sql"
	"errors"
)

type LoginRepositoryImpl struct {
}

func NewLoginRepository() LoginRepository {
	return &LoginRepositoryImpl{}
}

func (repository *LoginRepositoryImpl) Validate(ctx context.Context, db *sql.DB, login *domain.Login) error {

	SQL := "SELECT (username) FROM public.login WHERE username = ($1) AND password = ($2)"
	rows, err := db.QueryContext(ctx, SQL, login.Username, login.Password)

	if err != nil {
		return errors.New("invalid request")
	}

	defer rows.Close()

	if !rows.Next() {
		return errors.New("wrong username or password")
	}

	return nil

}

func (repository *LoginRepositoryImpl) CreateLogin(ctx context.Context, db *sql.DB, login *domain.Login) error {

	SQL := "INSERT INTO public.login (username, password) VALUES ($1, $2)"
	_, err := db.ExecContext(ctx, SQL, login.Username, login.Password)

	if err != nil {
		return errors.New("username already exist")
	}

	return nil
}
