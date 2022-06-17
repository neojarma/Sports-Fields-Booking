package logincontroller

import (
	"booking_fields/model/request"
	loginservice "booking_fields/service/login_service"
	"context"
	"net/http"

	"github.com/labstack/echo"
)

type LoginControllerImpl struct {
	LoginService loginservice.LoginService
}

func NewLoginController(loginService loginservice.LoginService) LoginController {
	return &LoginControllerImpl{
		LoginService: loginService,
	}
}

func (controller *LoginControllerImpl) Validate(ctx echo.Context) error {

	type Response struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
	}

	ctxBack := context.Background()

	req := new(request.LoginRequest)

	err := ctx.Bind(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &Response{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
	}

	err = controller.LoginService.Validate(ctxBack, req)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, &Response{
			Code:   http.StatusUnauthorized,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   200,
		Status: "success login",
	})
}
