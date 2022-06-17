package logincontroller

import (
	"github.com/labstack/echo"
)

type LoginController interface {
	Validate(ctx echo.Context) error
}
