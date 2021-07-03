package controllers

import (
	"app/configs"
	"app/models/rekening"
	"net/http"

	"github.com/labstack/echo"
)

// rekening

func CreateRekeningControllers(c echo.Context) error {
	var rekeningCreate rekening.RekeningRegister
	c.Bind(&rekeningCreate)
	var rekeningDB rekening.Rekening
	rekeningDB.Method_id = rekeningCreate.Method_id
	rekeningDB.Nama = rekeningCreate.Nama
	rekeningDB.Nomor_rekening = rekeningCreate.Nomor_rekening
	rekeningDB.Bank = rekeningCreate.Bank
	err := configs.DB.Create(&rekeningDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rekening.RekeningResponse{
			false, "failed insert rekening database", nil,
		})
	}

	return c.JSON(http.StatusOK, rekening.RekeningResponseSingle{
		true, "success insert rekening database", rekeningDB,
	})
}

func GetRekeningControllers(c echo.Context) error {
	var dataRekening []rekening.Rekening
	err := configs.DB.Find(&dataRekening).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rekening.RekeningResponse{
			false, "failed get rekening database", nil,
		})
	}

	return c.JSON(http.StatusOK, rekening.RekeningResponse{
		true, "success get rekening database", dataRekening,
	})
}

func GetRekeningByID(c echo.Context) error {
	var dataRekening []rekening.Rekening
	rekeningID := c.Param("rekeningid")
	err := configs.DB.First(&dataRekening, rekeningID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rekening.RekeningResponse{
			false, "failed get rekening database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, rekening.RekeningResponse{
		true, "success get rekening database by id", dataRekening,
	})
}

func UpdateRekeningControllers(c echo.Context) error {
	var rekeningUpdate rekening.RekeningUpdater
	c.Bind(&rekeningUpdate)
	var rekeningDB rekening.Rekening
	rekeningDB.Method_id = rekeningUpdate.Method_id
	rekeningDB.Nama = rekeningUpdate.Nama
	rekeningDB.Nomor_rekening = rekeningUpdate.Nomor_rekening
	rekeningDB.Bank = rekeningUpdate.Bank
	rekeningID := c.Param("rekeningid")

	err := configs.DB.Where("id = ?", rekeningID).Updates(&rekeningDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rekening.RekeningResponse{
			false, "failed update rekening database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, rekening.RekeningResponse{
		true, "success update rekening database by id", nil,
	})
}

func DeleteRekeningControllers(c echo.Context) error {
	var dataRekening []rekening.Rekening
	rekeningID := c.Param("rekeningid")
	err := configs.DB.Delete(&dataRekening, rekeningID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rekening.RekeningResponse{
			false, "failed delete rekening database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, rekening.RekeningResponse{
		true, "success delete rekening database by id", nil,
	})
}
