package crop

import (
	"net/http"
	"strconv"

	crop "api1-multi.com/a/src/Parcels/application/Crop"
	"api1-multi.com/a/src/Parcels/infrastructure"
	"github.com/gin-gonic/gin"
)

type DeleteCropC struct {
	uc crop.DeleteCropUC
}

func NewDeleteCropC()*DeleteCropC{

	mysql := infrastructure.GetCropMySQL()
	uc := crop.NewDeleteCropUC(mysql)

	return&DeleteCropC{uc: *uc}
}

func(controller *DeleteCropC)Execute(c *gin.Context){
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
			"Error": "Error al eliminar información del cultivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Información del cultivo eliminado",
	})
}