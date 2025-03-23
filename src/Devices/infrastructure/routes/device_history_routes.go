package routes

import (
	devicehistory_application "api1-multi.com/a/src/Devices/application/Device_history"
	"api1-multi.com/a/src/Devices/infrastructure"
	devicehistory_controllers "api1-multi.com/a/src/Devices/infrastructure/controllers/Device_history"
	"github.com/gin-gonic/gin"
)

func DeviceHistoryRoutes(r *gin.Engine){
	hm := infrastructure.NewDeviceHistoryMySQL()

	rhuc := devicehistory_application.NewRegisterDeviceInHistoryUC(hm)
	rhc := devicehistory_controllers.NewRegisterDeviceInHistoryC(*rhuc)

	ghuc := devicehistory_application.NewGetDeviceHistoryUC(hm)
	ghc := devicehistory_controllers.NewGetDeviceHistoryC(*ghuc)

	smduc := devicehistory_application.NewSetMaintenanceDateUC(hm)
	smdc := devicehistory_controllers.NewSetMaintenanceDateC(*smduc)

	srduc := devicehistory_application.NewSetRetireDateUC(hm)
	srdc := devicehistory_controllers.NewSetRetireDateC(*srduc)

	device_history := r.Group("device_history")
	{
		device_history.POST("/:id_parcel", rhc.Execute)
		device_history.GET("/:id_parcel", ghc.Execute)
		device_history.PATCH("/:id", smdc.Execute)
		device_history.PUT("/:id", srdc.Execute)
	}
}