package controllers

import (
	"net/http"
	"strconv"

	"api1-multi.com/a/src/Measurement/application"
	"github.com/gin-gonic/gin"
)

type GetOneMeasurementC struct {
	uc application.GetOneMeasurementUC
}

func NewGetOneMeasurementC(uc application.GetOneMeasurementUC)*GetOneMeasurementC{
	return&GetOneMeasurementC{uc: uc}
}

func(controller *GetOneMeasurementC)Execute(c *gin.Context){
	id, error_param := c.Params.Get("id_user")
	if !error_param {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "No se pudo mapear el parámetro",
		})
		return
	}

	id_number, error_strconv := strconv.Atoi(id)
	if error_strconv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Parámetro incorrecto",
		})
		return
	}

	results, err := controller.uc.Execute(id_number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener medición específica de la parcela",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Medición de la parcela obtenida",
		"Results": results,
	})
}