package routes

import (
	"api1-multi.com/a/src/Measurement/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func MeasurementRoutes(r *gin.Engine){
	measurement := r.Group("measurements")
	{
		measurement.POST("/", controllers.NewRegisterMeasurementC().Execute)
		measurement.GET("/:id_parcel", controllers.NewGetAllMeasurementsC().Execute)
		measurement.GET("/one/:id", controllers.NewGetOneMeasurementC().Execute)
	}
}