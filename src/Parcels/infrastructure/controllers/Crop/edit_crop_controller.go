package crop

import (
	"net/http"
	"strconv"

	crop "api1-multi.com/a/src/Parcels/application/Crop"
	"github.com/gin-gonic/gin"
)

type EditNameCropC struct {
	uc crop.EditNameCropUC
}

func NewEditNameCropC(uc crop.EditNameCropUC)*EditNameCropC{
	return&EditNameCropC{uc: uc}
}

func(controller *EditNameCropC)Execute(c *gin.Context){
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

	var input struct{
		Name string `json:"new_name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.uc.Execute(id_number, input.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al editar nombre del cultivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Nombre de cultivo editado",
	})
}