package params

import (
	"net/http"
	"strconv"

	params "api1-multi.com/a/src/Parcels/application/Params"
	"api1-multi.com/a/src/Parcels/infrastructure"
	"github.com/gin-gonic/gin"
)

type DeleteParamsC struct {
	uc params.DeleteParamsUC
}

func NewDeleteParamsC()*DeleteParamsC{

	mysql := infrastructure.GetParamsMySQL()
	uc := params.NewDeleteParamsUC(mysql)

	return&DeleteParamsC{uc: *uc}
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