package usercontroller

import (
	"booking_fields/model/request"
	"booking_fields/model/response"
	userservice "booking_fields/service/user_service"
	"context"
	"net/http"

	"github.com/labstack/echo"
)

type UserControllerImpl struct {
	UserService userservice.UserService
}

func NewUserController(service userservice.UserService) UserController {
	return &UserControllerImpl{
		UserService: service,
	}
}

type Response struct {
	Code   int                    `json:"code"`
	Status string                 `json:"status"`
	Data   *response.UserResponse `json:"data,omitempty"`
}

func (controller *UserControllerImpl) GetUserByUsername(ctx echo.Context) error {
	ctxBack := context.Background()
	username := ctx.Param("username")

	result, err := controller.UserService.GetUserByUsername(ctxBack, username)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, &Response{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "sucess get data",
		Data:   &result,
	})
}

func (controller *UserControllerImpl) UpdateUser(ctx echo.Context) error {
	ctxBack := context.Background()

	req := new(request.UserRequest)
	ctx.Bind(req)

	result, err := controller.UserService.UpdateUser(ctxBack, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   &result,
		})
	}

	return ctx.JSON(http.StatusOK, &Response{
		Code:   http.StatusOK,
		Status: "sucess update data",
		Data:   &result,
	})
}
