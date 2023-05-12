package controllers

import (
	"go_movie-ticket/lib/database"
	"go_movie-ticket/middlewares"
	"go_movie-ticket/models"
	"go_movie-ticket/models/payload"
	"go_movie-ticket/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Create user controller
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := database.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(400, "Failed to hash password")
	}

	user.Password = string(passwordHash)

	if err != nil {
		return nil
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

// Get user by id controller --
func GetUserByIdController(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	user, err := database.GetUserById(id)

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
	req := payload.UpdateUser{}

	id := middlewares.ExtractTokenUserId(c)

	c.Bind(&req)

	user, err := usecase.UpdateUser(id, req)
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
	id := middlewares.ExtractTokenUserId(c)

	_, err := database.DeleteUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success delete user",
	})
}

func TopupController(c echo.Context) error {
	req := payload.TopupRequest{}
	id := middlewares.ExtractTokenUserId(c)

	c.Bind(&req)

	user, err := usecase.TopupUser(uint(id), req.Saldo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success topup",
		"Data":    user,
	})
}
