package routes

import (
	"go_movie-ticket/constants"
	"go_movie-ticket/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// var jwtMiddleware = middleware.JWT([]byte(constants.SECRET_JWT))

func New() *echo.Echo {
	e := echo.New()

	// Auth
	e.POST("/register/users", controllers.CreateUserController)
	e.POST("/login/users", controllers.LoginUserController)

	// user
	user := e.Group("/profile", middleware.JWT([]byte(constants.SECRET_JWT)))
	user.GET("", controllers.GetUserByIdController)
	user.PUT("", controllers.UpdateUserByIdController)
	user.DELETE("", controllers.DeleteUserByIdController)

	// movie
	movie := e.Group("/movie")
	movie.GET("", controllers.GetMoviesControllers)
	movie.GET("/:id", controllers.GetMovieByIdController)

	// order
	e.POST("/order", controllers.CreateOrderController, middleware.JWT([]byte(constants.SECRET_JWT)))
	e.GET("/order", controllers.GetOrderByUserIdControllers, middleware.JWT([]byte(constants.SECRET_JWT)))

	return e
}
