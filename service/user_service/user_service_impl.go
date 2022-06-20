package userservice

import (
	"booking_fields/model/domain"
	"booking_fields/model/request"
	"booking_fields/model/response"
	userrepository "booking_fields/repository/user_repository"
	"context"
	"database/sql"
)

type UserServiceImpl struct {
	UserRepository userrepository.UserRepository
	Db             *sql.DB
}

func NewUserService(repo userrepository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: repo,
		Db:             db,
	}
}

func (service *UserServiceImpl) GetUserByUsername(ctx context.Context, username string) (response.UserResponse, error) {
	domainResult, err := service.UserRepository.GetUserByUsername(ctx, service.Db, username)

	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		IdUser:       domainResult.IdUser,
		Name:         domainResult.Name,
		Address:      domainResult.Address,
		PhoneNumber:  domainResult.PhoneNumber,
		Email:        domainResult.Email,
		ImageProfile: domainResult.ImageProfile,
	}, nil
}

func (service *UserServiceImpl) UpdateUser(ctx context.Context, request *request.UserRequest) (response.UserResponse, error) {

	newUsers := domain.User{
		IdUser:       request.IdUser,
		Name:         request.Name,
		Address:      request.Address,
		PhoneNumber:  request.PhoneNumber,
		Email:        request.Email,
		ImageProfile: request.ImageProfile,
	}

	err := service.UserRepository.UpdateUser(ctx, service.Db, &newUsers)

	if err != nil {
		return response.UserResponse{}, nil
	}

	return response.UserResponse{
		IdUser:       request.IdUser,
		Name:         request.Name,
		Address:      request.Address,
		PhoneNumber:  request.PhoneNumber,
		Email:        request.Email,
		ImageProfile: request.ImageProfile,
	}, nil
}
