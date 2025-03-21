package crop

import (
	"net/http"
	"strconv"

	crop "api1-multi.com/a/src/Parcels/application/Crop"
	"github.com/gin-gonic/gin"
)

type GetCropInfoC struct {
	uc crop.GetCropInfoUC
}

func NewGetCropInfoC(uc crop.GetCropInfoUC)*GetCropInfoC{
	return&GetCropInfoC{uc: uc}
}

func(controller *GetCropInfoC)Execute(c *gin.Context){
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
			"Error": "Error al obtener información de cultivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Información del cultivo obtenido",
		"Results": results,
	})
}