package params

import (
	"net/http"
	"strconv"

	params "api1-multi.com/a/src/Parcels/application/Params"
	"github.com/gin-gonic/gin"
)

type DeleteParamsC struct {
	uc params.DeleteParamsUC
}

func NewDeleteParamsC(uc params.DeleteParamsUC)*DeleteParamsC{
	return&DeleteParamsC{uc: uc}
}

func(controller *DeleteParamsC)Execute(c *gin.Context){
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
	
	if err := controller.uc.Execute(id_number); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al eliminar parámetros de monitoreo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Parámetros de cultivo eliminados",
	})
}