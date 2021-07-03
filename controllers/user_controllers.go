package controllers

import (
	"app/configs"
	"app/models/user"
	"net/http"

	"github.com/labstack/echo"
)

// user

func CreateUsersControllers(c echo.Context) error {
	var userCreate user.UserRegister
	c.Bind(&userCreate)
	var userDB user.User
	userDB.Name = userCreate.Name
	userDB.Email = userCreate.Email
	userDB.Password = userCreate.Password
	userDB.Mobile = userCreate.Mobile
	err := configs.DB.Create(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "failed insert user database", nil,
		})
	}

	return c.JSON(http.StatusOK, user.UserResponseSingle{
		true, "success insert user database", userDB,
	})
}

func GetUsersControllers(c echo.Context) error {
	var dataUsers []user.User
	err := configs.DB.Find(&dataUsers).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "failed get user database", nil,
		})
	}

	return c.JSON(http.StatusOK, user.UserResponse{
		true, "success get user database", dataUsers,
	})
}

func GetUserByID(c echo.Context) error {
	var dataUsers []user.User
	userID := c.Param("userid")
	err := configs.DB.First(&dataUsers, userID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "failed get user database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, user.UserResponse{
		true, "success get user database by id", dataUsers,
	})
}

func UpdateUserControllers(c echo.Context) error {
	var userUpdate user.UserUpdater
	c.Bind(&userUpdate)
	var userDB user.User
	userDB.Name = userUpdate.Name
	userDB.Email = userUpdate.Email
	userDB.Password = userUpdate.Password
	userDB.Mobile = userUpdate.Mobile
	userID := c.Param("userid")

	err := configs.DB.Where("id = ?", userID).Updates(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "failed update user database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, user.UserResponse{
		true, "success update user database by id", nil,
	})
}

func DeleteUserControllrs(c echo.Context) error {
	var dataUsers []user.User
	userID := c.Param("userid")
	err := configs.DB.Delete(&dataUsers, userID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user.UserResponse{
			false, "failed delete user database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, user.UserResponse{
		true, "success delete user database by id", nil,
	})
}
