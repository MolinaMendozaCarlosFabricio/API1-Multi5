package routes

import (
	crop_application "api1-multi.com/a/src/Parcels/application/Crop"
	"api1-multi.com/a/src/Parcels/infrastructure"
	crop_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Crop"
	"github.com/gin-gonic/gin"
)

func CropRoutes(r *gin.Engine){
	cm := infrastructure.NewCropMySQL()

	gciuc := crop_application.NewGetCropInfoUC(cm)
	gcic := crop_controllers.NewGetCropInfoC(*gciuc)

	ecnuc := crop_application.NewEditNameCropUC(cm)
	ecnc := crop_controllers.NewEditNameCropC(*ecnuc)

	dcuc := crop_application.NewDeleteCropUC(cm)
	dcc := crop_controllers.NewDeleteCropC(*dcuc)

	crop := r.Group("crop")
	{
		crop.GET("/:id", gcic.Execute)
		crop.PATCH("/:id", ecnc.Execute)
		crop.DELETE("/:id", dcc.Execute)
	}
}