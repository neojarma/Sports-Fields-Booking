package routes

import (
	logincontroller "booking_fields/controller/login_controller"
	registercontroller "booking_fields/controller/register_controller"
	reservationcontroller "booking_fields/controller/reservation_controller"
	usercontroller "booking_fields/controller/user_controller"
	venuecontroller "booking_fields/controller/venue_controller"
	loginrepository "booking_fields/repository/login_repository"
	registerrepository "booking_fields/repository/register_repository"
	reservationrepository "booking_fields/repository/reservation_repository"
	userrepository "booking_fields/repository/user_repository"
	venuerepository "booking_fields/repository/venue_repository"
	loginservice "booking_fields/service/login_service"
	registerservice "booking_fields/service/register_service"
	reservationservice "booking_fields/service/reservation_service"
	userservice "booking_fields/service/user_service"
	venueservice "booking_fields/service/venue_service"
	"database/sql"
	_ "embed"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func Routes(group *echo.Group, db *sql.DB) {
	reservationRoutes(group, db)
	UserRoutes(group, db)
	venueRoutes(group, db)
	registerRoutes(group, db)
	loginRoutes(group, db)
	swaggerRoute(group)
}

func swaggerRoute(group *echo.Group) {
	group.Static("/swagger", "dist")
}

func reservationRoutes(group *echo.Group, db *sql.DB) {
	repo := reservationrepository.NewReservationRepository()
	service := reservationservice.NewReservationService(repo, db)
	controller := reservationcontroller.NewReservationController(service)

	group.GET("/reserve", controller.GetReservationSchedule)
	group.GET("/reserve/schedule", controller.GetReservationScheduleForUpdate)
	group.GET("/reserve/:id", controller.GetUserReservationById)
	group.POST("/reserve", controller.CreateReservation)
	group.PUT("/reserve", controller.UpdateReservation)
	group.DELETE("/reserve/:id", controller.CancelReservation)
}

func UserRoutes(group *echo.Group, db *sql.DB) {
	repo := userrepository.NewUserRepository()
	service := userservice.NewUserService(repo, db)
	controller := usercontroller.NewUserController(service)

	group.GET("/user/:username", controller.GetUserByUsername)
	group.PUT("/user", controller.UpdateUser)

}

func venueRoutes(group *echo.Group, db *sql.DB) {
	repo := venuerepository.NewVenueRepository()
	service := venueservice.NewVenueService(repo, db)
	controller := venuecontroller.NewVenueController(service)

	group.GET("/venue", controller.GetAllVenue)
	group.GET("/venue/:id", controller.FindVenueById)

}

func registerRoutes(group *echo.Group, db *sql.DB) {
	repo := registerrepository.NewRegisterRepository(loginrepository.NewLoginRepository())
	service := registerservice.NewRegisterService(repo, db)
	controller := registercontroller.NewRegisterController(service)

	group.POST("/register", controller.Register)

}

func loginRoutes(group *echo.Group, db *sql.DB) {
	repo := loginrepository.NewLoginRepository()
	service := loginservice.NewLoginService(repo, db)
	controller := logincontroller.NewLoginController(service)

	group.POST("/login", controller.Validate)

}
