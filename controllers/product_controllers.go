package controllers

import (
	"app/configs"
	"app/models/product"
	"net/http"

	"github.com/labstack/echo"
)

// product

func CreateProductControllers(c echo.Context) error {
	var productCreate product.ProductRegister
	c.Bind(&productCreate)
	var productDB product.Product
	productDB.Category_id = productCreate.Category_id
	productDB.Kandang = productCreate.Kandang
	productDB.Jumlah = productCreate.Jumlah
	productDB.Hargamin = productCreate.Hargamin
	productDB.Hargamax = productCreate.Hargamax
	err := configs.DB.Create(&productDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, product.ProductResponse{
			false, "failed insert product database", nil,
		})
	}

	return c.JSON(http.StatusOK, product.ProductResponseSingle{
		true, "success insert product database", productDB,
	})
}

func GetProductControllers(c echo.Context) error {
	var dataProduct []product.Product
	err := configs.DB.Find(&dataProduct).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, product.ProductResponse{
			false, "failed get product database", nil,
		})
	}

	return c.JSON(http.StatusOK, product.ProductResponse{
		true, "success get product database", dataProduct,
	})
}

func GetProductByID(c echo.Context) error {
	var dataProduct []product.Product
	productID := c.Param("productid")
	err := configs.DB.First(&dataProduct, productID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, product.ProductResponse{
			false, "failed get product database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, product.ProductResponse{
		true, "success get product database by id", dataProduct,
	})
}

func UpdateProductControllers(c echo.Context) error {
	var productUpdate product.ProductUpdater
	c.Bind(&productUpdate)
	var productDB product.Product
	productDB.Category_id = productUpdate.Category_id
	productDB.Kandang = productUpdate.Kandang
	productDB.Jumlah = productUpdate.Jumlah
	productDB.Hargamin = productUpdate.Hargamin
	productDB.Hargamax = productUpdate.Hargamax
	productID := c.Param("productid")

	err := configs.DB.Where("id = ?", productID).Updates(&productDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, product.ProductResponse{
			false, "failed update product database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, product.ProductResponse{
		true, "success update product database by id", nil,
	})
}

func DeleteProductControllers(c echo.Context) error {
	var dataProduct []product.Product
	productID := c.Param("productid")
	err := configs.DB.Delete(&dataProduct, productID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, product.ProductResponse{
			false, "failed delete product database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, product.ProductResponse{
		true, "success delete product database by id", nil,
	})
}
