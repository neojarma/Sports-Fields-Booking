package registercontroller

import "github.com/labstack/echo"

type RegisterController interface {
	Register(ctx echo.Context) error
}
