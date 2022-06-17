package loginservice

import (
	"booking_fields/model/request"
	"context"
)

type LoginService interface {
	Validate(ctx context.Context, request *request.LoginRequest) error
	CreateLogin(ctx context.Context, request *request.LoginRequest) error
}
