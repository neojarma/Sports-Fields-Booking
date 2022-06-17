package main

import (
	"booking_fields/connection"
	"booking_fields/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	db, err := connection.GetConnection()
	if err != nil {
		panic(err)
	}

	group := e.Group("/api")
	routes.Routes(group, db)

	e.Logger.Fatal(e.Start(":8080"))
}
