package croptype

import (
	"net/http"
	"strconv"

	croptype "api1-multi.com/a/src/Parcels/application/Crop_type"
	"api1-multi.com/a/src/Parcels/infrastructure"
	"github.com/gin-gonic/gin"
)

type GetCropTypeC struct {
	uc croptype.GetCropTypeUC
}

func NewGetCropTypeC()*GetCropTypeC{

	mysql := infrastructure.GetCropSTypeMySQL()
	uc := croptype.NewGetCropTypeUC(mysql)

	return&GetCropTypeC{uc: *uc}
}

func(controller *GetCropTypeC)Execute(c *gin.Context){
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
			"Error": "Error al obtener tipo de cultivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Tipo de cultivo obtenidos",
		"Results": results,
	})
}