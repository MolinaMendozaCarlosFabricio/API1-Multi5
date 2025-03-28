package devicehistory

import (
	"net/http"
	"strconv"

	devicehistory "api1-multi.com/a/src/Devices/application/Device_history"
	"api1-multi.com/a/src/Devices/infrastructure"
	"github.com/gin-gonic/gin"
)

type SetMaintenanceDateC struct {
	uc devicehistory.SetMaintenanceDateUC
}

func NewSetMaintenanceDateC()*SetMaintenanceDateC{

	mysql := infrastructure.GetDeviceHistoryMySQL()
	uc := devicehistory.NewSetMaintenanceDateUC(mysql)

	return&SetMaintenanceDateC{uc: *uc}
}

func(controller *SetMaintenanceDateC)Execute(c *gin.Context){
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
			"Error": "Error al añadir fecha de mantenimiento del dispositivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Fecha de mantenimiento de dispositivo actualizado",
	})
}