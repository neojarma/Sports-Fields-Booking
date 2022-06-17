package registerrepository

import (
	"booking_fields/model/domain"
	loginrepository "booking_fields/repository/login_repository"
	"context"
	"database/sql"
)

type RegisterRepositoryImpl struct {
	loginRepo loginrepository.LoginRepository
}

func NewRegisterRepository(loginRepository loginrepository.LoginRepository) RegisterRepository {
	return &RegisterRepositoryImpl{loginRepo: loginRepository}
}

func (repository *RegisterRepositoryImpl) Register(ctx context.Context, db *sql.DB, register *domain.Register) error {

	request := domain.Login{
		Username: register.Username,
		Password: register.Password,
	}

	// add loginrequest to db
	err := repository.loginRepo.CreateLogin(ctx, db, &request)
	if err != nil {
		return err
	}

	// add new register data to db
	SQL := "INSERT INTO public.user (id_user, name, address, phone_number, email, username) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = db.ExecContext(ctx, SQL, register.IdUser, register.Name, register.Address, register.PhoneNumber, register.Email, register.Username)

	if err != nil {
		return err
	}

	return nil
}
