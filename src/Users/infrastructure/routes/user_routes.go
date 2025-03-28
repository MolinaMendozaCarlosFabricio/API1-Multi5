package routes

import (
	"api1-multi.com/a/src/Users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){
	user := r.Group("users")
	{
		user.POST("/", controllers.NewRegisterUserController().Execute)
		user.GET("/:id", controllers.NewViewUserController().Execute)
		user.PUT("/user/:id", controllers.NewEditUserController().Execute)
		user.PUT("/credential/:id", controllers.NewEditCredentialsController().Execute)
		user.DELETE("/:id", controllers.NewDeleteUserController().Execute)
	}
}