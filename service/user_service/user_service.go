package userservice

import (
	"booking_fields/model/request"
	"booking_fields/model/response"
	"context"
)

type UserService interface {
	GetUserByUsername(ctx context.Context, username string) (response.UserResponse, error)
	UpdateUser(ctx context.Context, request *request.UserRequest) (response.UserResponse, error)
}
