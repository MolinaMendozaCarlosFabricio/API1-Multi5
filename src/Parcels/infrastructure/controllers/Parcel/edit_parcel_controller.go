package parcel

import (
	"net/http"
	"strconv"

	parcel "api1-multi.com/a/src/Parcels/application/Parcel"
	"api1-multi.com/a/src/Parcels/infrastructure"
	"github.com/gin-gonic/gin"
)

type EditParcelC struct {
	uc parcel.EditParcelUC
}

func NewEditParcelC()*EditParcelC{

	mysql := infrastructure.GetParcelMysQL()
	uc := parcel.NewEditParcelUC(mysql)

	return&EditParcelC{uc: *uc}
}

func(controller *EditParcelC)Execute(c *gin.Context){
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
		Id_device int `json:"id_device"`
		Id_status int `json:"id_status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.uc.Execute(
		id_number,
		input.Id_device, 
		input.Id_status,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al editar parcela",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Parcela editada",
	})

}