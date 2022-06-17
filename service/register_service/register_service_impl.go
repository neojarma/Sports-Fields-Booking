package registerservice

import (
	"booking_fields/helper"
	"booking_fields/model/domain"
	"booking_fields/model/request"
	registerrepository "booking_fields/repository/register_repository"
	"context"
	"database/sql"
	"strconv"
)

type RegisterServiceImpl struct {
	RegisterRepository registerrepository.RegisterRepository

	Db *sql.DB
}

func NewRegisterService(repo registerrepository.RegisterRepository, db *sql.DB) RegisterService {
	return &RegisterServiceImpl{
		RegisterRepository: repo,
		Db:                 db,
	}
}

func (service *RegisterServiceImpl) Register(ctx context.Context, request *request.RegisterRequest) error {

	idUser := "USER-" + strconv.Itoa(helper.GenerateRandomId())

	regist := domain.Register{
		IdUser:      idUser,
		Name:        request.Name,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Username:    request.Username,
		Password:    request.Password,
	}

	err := service.RegisterRepository.Register(ctx, service.Db, &regist)

	if err != nil {
		return err
	}

	return nil
}
