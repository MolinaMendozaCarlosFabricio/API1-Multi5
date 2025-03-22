package routes

import (
	device_application "api1-multi.com/a/src/Devices/application/Device"
	"api1-multi.com/a/src/Devices/infrastructure"
	device_controllers "api1-multi.com/a/src/Devices/infrastructure/controllers/Device"
	"github.com/gin-gonic/gin"
)

func DeviceRoutes(r *gin.Engine){
	dm := infrastructure.NewDeviceMySQL()

	gduc := device_application.NewGetDeviceUC(dm)
	gdc := device_controllers.NewGetDeviceC(*gduc)

	devices := r.Group("devices")
	{
		devices.GET("/:id", gdc.Execute)
	}
}