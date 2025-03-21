package parcel

import (
	"net/http"
	"strconv"

	parcel "api1-multi.com/a/src/Parcels/application/Parcel"
	"github.com/gin-gonic/gin"
)

type DeleteParcelC struct {
	uc parcel.DeleteParcelUC
}

func NewDeleteParcelC(uc parcel.DeleteParcelUC)*DeleteParcelC{
	return&DeleteParcelC{uc: uc}
}

func(controller *DeleteParcelC)Execute(c *gin.Context){
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
			"Error": "Error al eliminar parcela",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Parcela eliminada",
	})
}