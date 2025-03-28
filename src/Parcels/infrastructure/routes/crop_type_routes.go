package routes

import (
	croptype_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Crop_type"
	"github.com/gin-gonic/gin"
)

func CroptTypeRoutes(r *gin.Engine){
	crop_type := r.Group("crop_type")
	{
		crop_type.GET("/:id", croptype_controllers.NewGetCropTypeC().Execute)
	}
}