package loginservice

import (
	"booking_fields/model/domain"
	"booking_fields/model/request"
	loginrepository "booking_fields/repository/login_repository"
	"context"
	"database/sql"
)

type LoginServiceImpl struct {
	LoginRepository loginrepository.LoginRepository
	Db              *sql.DB
}

func NewLoginService(loginRepository loginrepository.LoginRepository, db *sql.DB) LoginService {
	return &LoginServiceImpl{
		LoginRepository: loginRepository,
		Db:              db,
	}
}

func (service *LoginServiceImpl) Validate(ctx context.Context, request *request.LoginRequest) error {
	login := domain.Login{
		Username: request.Username,
		Password: request.Password,
	}

	err := service.LoginRepository.Validate(ctx, service.Db, &login)

	if err != nil {
		return err
	}

	return nil
}

func (service *LoginServiceImpl) CreateLogin(ctx context.Context, request *request.LoginRequest) error {
	login := domain.Login{
		Username: request.Username,
		Password: request.Password,
	}

	err := service.LoginRepository.CreateLogin(ctx, service.Db, &login)

	if err != nil {
		return err
	}

	return nil
}
