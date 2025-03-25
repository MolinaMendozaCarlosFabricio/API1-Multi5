package routes

import (
	crop "api1-multi.com/a/src/Parcels/application/Crop"
	params "api1-multi.com/a/src/Parcels/application/Params"
	parcel_application "api1-multi.com/a/src/Parcels/application/Parcel"
	"api1-multi.com/a/src/Parcels/infrastructure"
	parcel_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Parcel"
	"github.com/gin-gonic/gin"
)

func ParcelRoutes(r *gin.Engine){
	pm := infrastructure.NewParcelMySQL()
	pcm := infrastructure.NewParamsMySQL()
	cm := infrastructure.NewCropMySQL()

	cpuc := parcel_application.NewCreateParcelUC(pm)
	cpcuc := params.NewSetParamsUC(pcm)
	ccuc := crop.NewCreateCropUC(cm)
	cpc := parcel_controllers.NewCreateParcelC(*ccuc, *cpcuc, *cpuc)

	gmpuc := parcel_application.NewGetAllMyParcelsUC(pm)
	gmpc := parcel_controllers.NewGetAllMyParcelsC(*gmpuc)

	gopuc := parcel_application.NewGetAllAboutMyParcel(pm)
	gopc := parcel_controllers.NewGetAllAboutMyParcelC(*gopuc)

	epuc := parcel_application.NewEditParcelUC(pm)
	epc := parcel_controllers.NewEditParcelC(*epuc)

	dpuc := parcel_application.NewDeleteParcelUC(pm)
	dpc := parcel_controllers.NewDeleteParcelC(*dpuc)

	parcels := r.Group("parcels")
	{
		parcels.POST("/", cpc.Execute)
		parcels.GET("/:id_user", gmpc.Execute)
		parcels.GET("/all_about/:id", gopc.Execute)
		parcels.PUT("/:id", epc.Execute)
		parcels.DELETE("/:id", dpc.Execute)
	}
}