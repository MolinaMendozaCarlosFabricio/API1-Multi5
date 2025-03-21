package params

import (
	"net/http"
	"strconv"

	params "api1-multi.com/a/src/Parcels/application/Params"
	"github.com/gin-gonic/gin"
)

type GetParamsC struct {
	uc params.GetParamsUC
}

func NewGetParamsC(uc params.GetParamsUC)*GetParamsC{
	return&GetParamsC{uc: uc}
}

func(controller *GetParamsC)Execute(c *gin.Context){
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

	results, err := controller.uc.Execute(id_number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener parámetros de monitoreo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Parámetros de cultivo obtenidos",
		"Results": results,
	})
}