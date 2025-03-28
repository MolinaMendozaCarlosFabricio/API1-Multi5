package routes

import (
	parcel_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Parcel"
	"github.com/gin-gonic/gin"
)

func ParcelRoutes(r *gin.Engine){
	parcels := r.Group("parcels")
	{
		parcels.POST("/", parcel_controllers.NewCreateParcelC().Execute)
		parcels.GET("/:id_user", parcel_controllers.NewGetAllMyParcelsC().Execute)
		parcels.GET("/all_about/:id", parcel_controllers.NewGetAllAboutMyParcelC().Execute)
		parcels.PUT("/:id", parcel_controllers.NewEditParcelC().Execute)
		parcels.DELETE("/:id", parcel_controllers.NewDeleteParcelC().Execute)
	}
}