package routes

import (
	params_controllers "api1-multi.com/a/src/Parcels/infrastructure/controllers/Params"
	"github.com/gin-gonic/gin"
)

func ParamsRoutes(r *gin.Engine){
	params := r.Group("parameters_cultivation")
	{
		params.GET("/:id", params_controllers.NewGetParamsC().Execute)
		params.PUT("/:id", params_controllers.NewEditParametersC().Execute)
		params.DELETE("/:id", params_controllers.NewDeleteParamsC().Execute)
	}
}