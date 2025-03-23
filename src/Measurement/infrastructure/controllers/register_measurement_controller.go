package controllers

import (
	"net/http"

	"api1-multi.com/a/src/Measurement/application"
	"github.com/gin-gonic/gin"
)

type RegisterMeasurementC struct {
	uc application.RegisterMeasurementUC
}

func NewRegisterMeasurementC(uc application.RegisterMeasurementUC)*RegisterMeasurementC{
	return&RegisterMeasurementC{uc: uc}
}

func(controller *RegisterMeasurementC)Execute(c *gin.Context){
	var input struct{
		Id_parcel int `json:"id_parcel"`
		Temp float32 `json:"temp"`
		Humedity float32 `json:"humedity"`
		Air float32 `json:"air"`
		Sun float32 `json:"sun"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
			"Err": err,
		})
		return
	}

	if err := controller.uc.Execute(
		input.Id_parcel,
		input.Temp,
		input.Humedity,
		input.Air,
		input.Sun,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar la medición",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Medición registrada",
	})
	
}