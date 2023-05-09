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
	// user
	e.POST("/register/users", controllers.CreateUserController)
	e.POST("/login/users", controllers.LoginUserController)
	user := e.Group("/user", middleware.JWT([]byte(constants.SECRET_JWT)))
	user.GET("", controllers.GetUsersController)
	user.GET("/:id", controllers.GetUserByIdController)
	user.PUT("/:id", controllers.UpdateUserByIdController)
	user.DELETE("/:id", controllers.DeleteUserByIdController)

	// movie
	e.POST("/movies", controllers.CreateMovieController)
	e.GET("/movies", controllers.GetMoviesControllers)
	e.GET("/movies/:id", controllers.GetMovieByIdController)
	e.PUT("/movies/:id", controllers.UpdateMovieByIdController)
	e.DELETE("/movies/:id", controllers.DeleteMovieByIdController)

	// order
	e.POST("/order", controllers.CreateOrderController)
	e.GET("/order", controllers.GetOrderControllers)
	// e.GET("/movies/:id", controllers.GetMovieByIdController)
	// e.PUT("/movies/:id", controllers.UpdateMovieByIdController)
	// e.DELETE("/movies/:id", controllers.DeleteMovieByIdController)
	
	return e
}
