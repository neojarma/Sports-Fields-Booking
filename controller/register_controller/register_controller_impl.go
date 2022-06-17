package registercontroller

import (
	"booking_fields/model/request"
	registerservice "booking_fields/service/register_service"
	"context"
	"net/http"

	"github.com/labstack/echo"
)

type RegisterControllerImpl struct {
	service registerservice.RegisterService
}

func NewRegisterController(service registerservice.RegisterService) RegisterController {
	return &RegisterControllerImpl{
		service: service,
	}
}

func (controller *RegisterControllerImpl) Register(ctx echo.Context) error {

	type Response struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
	}

	ctxBack := context.Background()

	req := new(request.RegisterRequest)
	err := ctx.Bind(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
	}

	err = controller.service.Register(ctxBack, req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "succes register",
	})
}
