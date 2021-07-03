package controllers

import (
	"app/configs"
	"app/models/category"
	"net/http"

	"github.com/labstack/echo"
)

// category

func CreateCategoryControllers(c echo.Context) error {
	var categoryCreate category.CategoryRegister
	c.Bind(&categoryCreate)
	var categoryDB category.Category
	categoryDB.Name = categoryCreate.Name
	err := configs.DB.Create(&categoryDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false, "failed insert category database", nil,
		})
	}

	return c.JSON(http.StatusOK, category.CategoryResponseSingle{
		true, "success insert category database", categoryDB,
	})
}

func GetCategoryControllers(c echo.Context) error {
	var dataCategorys []category.Category
	err := configs.DB.Find(&dataCategorys).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false, "failed get category database", nil,
		})
	}

	return c.JSON(http.StatusOK, category.CategoryResponse{
		true, "success get category database", dataCategorys,
	})
}

func GetCategoryByID(c echo.Context) error {
	var dataCategorys []category.Category
	categoryID := c.Param("categoryid")
	err := configs.DB.First(&dataCategorys, categoryID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false, "failed get category database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, category.CategoryResponse{
		true, "success get category database by id", dataCategorys,
	})
}

func UpdateCategoryControllers(c echo.Context) error {
	var categoryUpdate category.CategoryUpdater
	c.Bind(&categoryUpdate)
	var categoryDB category.Category
	categoryDB.Name = categoryUpdate.Name
	categoryID := c.Param("categoryid")

	err := configs.DB.Where("id = ?", categoryID).Updates(&categoryDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false, "failed update category database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, category.CategoryResponse{
		true, "success update category database by id", nil,
	})
}

func DeleteCategoryControllers(c echo.Context) error {
	var dataCategorys []category.Category
	categoryID := c.Param("categoryid")
	err := configs.DB.Delete(&dataCategorys, categoryID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, category.CategoryResponse{
			false, "failed delete category database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, category.CategoryResponse{
		true, "success delete category database by id", nil,
	})
}
