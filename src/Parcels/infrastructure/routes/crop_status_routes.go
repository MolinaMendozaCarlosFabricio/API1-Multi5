package routes

import (
	cropstatus_application "api1-multi.com/a/src/Parcels/application/Crop_status"
	"api1-multi.com/a/src/Parcels/infrastructure"
	cropstatus_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Crop_status"
	"github.com/gin-gonic/gin"
)

func CropStatusRoutes(r *gin.Engine){
	csm := infrastructure.NewCropStatusMySQL()

	gcsuc := cropstatus_application.NewGetCropStatusUC(csm)
	gcsc := cropstatus_controllers.NewGetCropStatusC(*gcsuc)

	crop_status := r.Group("crop_status")
	{
		crop_status.GET("/:id", gcsc.Execute)
	}
}