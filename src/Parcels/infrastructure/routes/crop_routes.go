package routes

import (
	crop_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Crop"
	"github.com/gin-gonic/gin"
)

func CropRoutes(r *gin.Engine){
	crop := r.Group("crop")
	{
		crop.GET("/:id", crop_controllers.NewGetCropInfoC().Execute)
		crop.PATCH("/:id", crop_controllers.NewEditNameCropC().Execute)
		crop.DELETE("/:id", crop_controllers.NewDeleteCropC().Execute)
	}
}