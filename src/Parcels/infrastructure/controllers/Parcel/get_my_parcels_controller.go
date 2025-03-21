package parcel

import (
	"net/http"
	"strconv"

	parcel "api1-multi.com/a/src/Parcels/application/Parcel"
	"github.com/gin-gonic/gin"
)

type GetAllMyParcelsC struct {
	uc parcel.GetAllMyParcelsUC
}

func NewGetAllMyParcelsC(uc parcel.GetAllMyParcelsUC)*GetAllMyParcelsC{
	return&GetAllMyParcelsC{uc: uc}
}

func(controller *GetAllMyParcelsC)Execute(c *gin.Context){
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
			"Error": "Error al obtener información de las parcelas",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Información del parcelas obtenidas",
		"Results": results,
	})
}