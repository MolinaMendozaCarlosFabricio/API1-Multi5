package routes

import (
	devicehistory_controllers "api1-multi.com/a/src/Devices/infrastructure/controllers/Device_history"
	"github.com/gin-gonic/gin"
)

func DeviceHistoryRoutes(r *gin.Engine){
	device_history := r.Group("device_history")
	{
		device_history.POST("/:id_parcel", devicehistory_controllers.NewRegisterDeviceInHistoryC().Execute)
		device_history.GET("/:id_parcel", devicehistory_controllers.NewGetDeviceHistoryC().Execute)
		device_history.PATCH("/:id", devicehistory_controllers.NewSetMaintenanceDateC().Execute)
		device_history.PUT("/:id", devicehistory_controllers.NewSetRetireDateC().Execute)
	}
}