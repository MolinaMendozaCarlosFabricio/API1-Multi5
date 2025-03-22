package device

import (
	"net/http"
	"strconv"

	device "api1-multi.com/a/src/Devices/application/Device"
	"github.com/gin-gonic/gin"
)

type GetDeviceC struct {
	uc device.GetDeviceUC
}

func NewGetDeviceC(uc device.GetDeviceUC)*GetDeviceC{
	return&GetDeviceC{uc: uc}
}

func(controller *GetDeviceC)Execute(c *gin.Context){
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
			"Error": "Error al obtener dispositivo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Dispositivo obtenido",
		"Results": results,
	})
}