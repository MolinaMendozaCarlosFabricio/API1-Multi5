package routes

import (
	cropstatus_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Crop_status"
	"github.com/gin-gonic/gin"
)

func CropStatusRoutes(r *gin.Engine){
	crop_status := r.Group("crop_status")
	{
		crop_status.GET("/:id", cropstatus_controllers.NewGetCropStatusC().Execute)
	}
}