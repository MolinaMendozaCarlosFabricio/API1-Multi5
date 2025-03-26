package controllers

import (
	"net/http"
	"strconv"

	"api1-multi.com/a/src/Measurement/application"
	"api1-multi.com/a/src/Measurement/infrastructure"
	"github.com/gin-gonic/gin"
)

type GetAllMeasurementsC struct {
	uc application.GetAllMeasurementsUC
}

func NewGetAllMeasurementsC()*GetAllMeasurementsC{

	mysql := infrastructure.GetMeasurementMySQL()
	uc := application.NewGetAllMeasurementsUC(mysql)

	return&GetAllMeasurementsC{uc: *uc}
}

func(controller *GetAllMeasurementsC)Execute(c *gin.Context){
	id, error_param := c.Params.Get("id_parcel")
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
			"Error": "Error al obtener mediciones de las parcelas",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Mediciones de las parcelas obtenidas",
		"Results": results,
	})
}