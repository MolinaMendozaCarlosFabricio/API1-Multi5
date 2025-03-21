package routes

import (
	params_application "api1-multi.com/a/src/Parcels/application/Params"
	"api1-multi.com/a/src/Parcels/infrastructure"
	params_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Params"
	"github.com/gin-gonic/gin"
)

func ParamsRoutes(r *gin.Engine){
	pm := infrastructure.NewParamsMySQL()

	gpuc := params_application.NewGetParamsUC(pm)
	gpc := params_controllers.NewGetParamsC(*gpuc)

	epuc := params_application.NewEditParametersUc(pm)
	epc := params_controllers.NewEditParametersC(*epuc)

	dpuc := params_application.NewDeleteParamsUC(pm)
	dpc := params_controllers.NewDeleteParamsC(*dpuc)

	params := r.Group("parameters_cultivation")
	{
		params.GET("/:id", gpc.Execute)
		params.PUT("/:id", epc.Execute)
		params.DELETE("/:id", dpc.Execute)
	}
}