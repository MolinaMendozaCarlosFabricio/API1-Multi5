package routes

import (
	"api1-multi.com/a/src/Measurement/application"
	"api1-multi.com/a/src/Measurement/infrastructure"
	"api1-multi.com/a/src/Measurement/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func MeasurementRoutes(r *gin.Engine){
	ms := infrastructure.NewMeasurementMySQL()

	rmuc := application.NewRegisterMeasurementUC(ms)
	rmc := controllers.NewRegisterMeasurementC(*rmuc)

	gamuc := application.NewGetAllMeasurementsUC(ms)
	gamc := controllers.NewGetAllMeasurementsC(*gamuc)

	gomun := application.NewGetOneMeasurementUC(ms)
	gomc := controllers.NewGetOneMeasurementC(*gomun)

	measurement := r.Group("measurements")
	{
		measurement.POST("/", rmc.Execute)
		measurement.GET("/:id_parcel", gamc.Execute)
		measurement.GET("/one/:id", gomc.Execute)
	}
}