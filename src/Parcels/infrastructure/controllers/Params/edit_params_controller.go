package params

import (
	"net/http"
	"strconv"

	params "api1-multi.com/a/src/Parcels/application/Params"
	"api1-multi.com/a/src/Parcels/domain/models"
	"github.com/gin-gonic/gin"
)

type EditParametersC struct {
	uc params.EditParametersUc
}

func NewEditParametersC(uc params.EditParametersUc)*EditParametersC{
	return&EditParametersC{uc: uc}
}

func(controller *EditParametersC)Execute(c *gin.Context){
	id, error_param := c.Params.Get("id")
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

	var input models.CultivationParameters

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.uc.Execute(
		id_number, 
		input.Humidity_min, 
		input.Humidity_max,
		input.Temp_min,
		input.Temp_max,
		input.Min_air_con,
		input.Max_air_con,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al editar parámetros de cultivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Parámetros de cultivo editados",
	})
}