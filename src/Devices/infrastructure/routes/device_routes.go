package routes

import (
	device_controllers "api1-multi.com/a/src/Devices/infrastructure/controllers/Device"
	"github.com/gin-gonic/gin"
)

func DeviceRoutes(r *gin.Engine){
	devices := r.Group("devices")
	{
		devices.GET("/:id", device_controllers.NewGetDeviceC().Execute)
	}
}