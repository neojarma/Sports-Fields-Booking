package usercontroller

import (
	"github.com/labstack/echo"
)

type UserController interface {
	GetUserByUsername(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
}
