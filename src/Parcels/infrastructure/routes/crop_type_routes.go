package routes

import (
	croptype_application "api1-multi.com/a/src/Parcels/application/Crop_type"
	"api1-multi.com/a/src/Parcels/infrastructure"
	croptype_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Crop_type"
	"github.com/gin-gonic/gin"
)

func CroptTypeRoutes(r *gin.Engine){
	ctm := infrastructure.NewCropTypeMySQL()

	gctuc := croptype_application.NewGetCropTypeUC(ctm)
	gctc := croptype_controllers.NewGetCropTypeC(*gctuc)

	crop_type := r.Group("crop_type")
	{
		crop_type.GET("/:id", gctc.Execute)
	}
}