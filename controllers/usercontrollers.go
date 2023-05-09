package controllers

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create user controller
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := database.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Data":    user,
		"Message": "success create user",
	})
}

// Login user
func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

// Get all user controller
func GetUsersController(c echo.Context) error {
	// err := middlewares.ExtractTokenUserId()
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	users, err := database.GetUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get all user",
		"Data":    users,
	})
}

// Get user by id controller
func GetUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	user, err := database.GetUserById(UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get user",
		"Data":    user,
	})
}

// Update user by id controller
func UpdateUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	user := models.User{}
	c.Bind(&user)

	user, err := database.UpdateUser(user, UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success update user",
		"Data":    user,
	})
}

// Delete user by id controller-
func DeleteUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	_, err := database.DeleteUser(UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success delete user",
	})
}
