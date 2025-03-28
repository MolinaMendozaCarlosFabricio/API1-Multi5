package devicehistory

import (
	"net/http"
	"strconv"

	devicehistory "api1-multi.com/a/src/Devices/application/Device_history"
	"api1-multi.com/a/src/Devices/infrastructure"
	"github.com/gin-gonic/gin"
)

type GetDeviceHistoryC struct {
	uc devicehistory.GetDeviceHistoryUC
}

func NewGetDeviceHistoryC()*GetDeviceHistoryC{

	mysql := infrastructure.GetDeviceHistoryMySQL()
	uc := devicehistory.NewGetDeviceHistoryUC(mysql)

	return&GetDeviceHistoryC{uc: *uc}
}

func(controller *GetDeviceHistoryC)Execute(c *gin.Context){
	id, error_param := c.Params.Get("id_parcel")
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
			"Error": "Error al obtener historial de dispositivos de la parcela",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Historial de dispositivos obtenido",
		"Results": results,
	})
}