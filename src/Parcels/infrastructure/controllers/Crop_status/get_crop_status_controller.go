package cropstatus

import (
	"net/http"
	"strconv"

	cropstatus "api1-multi.com/a/src/Parcels/application/Crop_status"
	"github.com/gin-gonic/gin"
)

type GetCropStatusC struct {
	uc cropstatus.GetCropStatusUC
}

func NewGetCropStatusC(uc cropstatus.GetCropStatusUC)*GetCropStatusC{
	return&GetCropStatusC{uc: uc}
}

func(controller *GetCropStatusC)Execute(c *gin.Context){
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
			"Error": "Error al obtener el estatus de cultivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Estatus de cultivo obtenidos",
		"Results": results,
	})
}