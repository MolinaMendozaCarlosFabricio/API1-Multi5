package parcel

import (
	"net/http"
	"strconv"

	parcel "api1-multi.com/a/src/Parcels/application/Parcel"
	"api1-multi.com/a/src/Parcels/infrastructure"
	"github.com/gin-gonic/gin"
)

type GetAllAboutMyParcelC struct {
	uc parcel.GetAllAboutMyParcelUC
}

func NewGetAllAboutMyParcelC()*GetAllAboutMyParcelC{

	mysql := infrastructure.GetParcelMysQL()
	uc := parcel.NewGetAllAboutMyParcel(mysql)

	return&GetAllAboutMyParcelC{uc: *uc}
}

func(controller *GetAllAboutMyParcelC)Execute(c *gin.Context){
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
			"Error": "Error al obtener información de la parcela",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Información de la parcela obtenida",
		"Results": results,
	})
}