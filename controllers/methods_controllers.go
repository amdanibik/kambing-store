package controllers

import (
	"app/configs"
	"app/models/methods"
	"net/http"

	"github.com/labstack/echo"
)

// methods

func CreateMethodsControllers(c echo.Context) error {
	var methodsCreate methods.MethodsRegister
	c.Bind(&methodsCreate)
	var methodsDB methods.Methods
	methodsDB.Name = methodsCreate.Name
	err := configs.DB.Create(&methodsDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, methods.MethodsResponse{
			false, "failed insert methods database", nil,
		})
	}

	return c.JSON(http.StatusOK, methods.MethodsResponseSingle{
		true, "success insert methods database", methodsDB,
	})
}

func GetMethodsControllers(c echo.Context) error {
	var dataMethodss []methods.Methods
	err := configs.DB.Find(&dataMethodss).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, methods.MethodsResponse{
			false, "failed get methods database", nil,
		})
	}

	return c.JSON(http.StatusOK, methods.MethodsResponse{
		true, "success get methods database", dataMethodss,
	})
}

func GetMethodsByID(c echo.Context) error {
	var dataMethodss []methods.Methods
	methodsID := c.Param("methodsid")
	err := configs.DB.First(&dataMethodss, methodsID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, methods.MethodsResponse{
			false, "failed get methods database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, methods.MethodsResponse{
		true, "success get methods database by id", dataMethodss,
	})
}

func UpdateMethodsControllers(c echo.Context) error {
	var methodsUpdate methods.MethodsUpdater
	c.Bind(&methodsUpdate)
	var methodsDB methods.Methods
	methodsDB.Name = methodsUpdate.Name
	methodsID := c.Param("methodsid")

	err := configs.DB.Where("id = ?", methodsID).Updates(&methodsDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, methods.MethodsResponse{
			false, "failed update methods database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, methods.MethodsResponse{
		true, "success update methods database by id", nil,
	})
}

func DeleteMethodsControllers(c echo.Context) error {
	var dataMethodss []methods.Methods
	methodsID := c.Param("methodsid")
	err := configs.DB.Delete(&dataMethodss, methodsID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, methods.MethodsResponse{
			false, "failed delete methods database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, methods.MethodsResponse{
		true, "success delete methods database by id", nil,
	})
}
