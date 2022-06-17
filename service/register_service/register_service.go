package registerservice

import (
	"booking_fields/model/request"
	"context"
)

type RegisterService interface {
	Register(ctx context.Context, request *request.RegisterRequest) error
}
