package routes

import (
	"api1-multi.com/a/src/Users/application"
	"api1-multi.com/a/src/Users/infrastructure"
	"api1-multi.com/a/src/Users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){

	db := infrastructure.NewUserMySQL()
	ru := application.NewRegisterUserUC(db)
	ruc := controllers.NewRegisterUserController(*ru)
	gou := application.NewViewUserUC(db)
	gouc := controllers.NewViewUserController(*gou)
	eu := application.NewEditUserUC(db)
	euc := controllers.NewEditUserController(*eu)
	ec := application.NewEditCredentialsUC(db)
	ecc := controllers.NewEditCredentialsController(*ec)
	du := application.NewDeleteUserUC(db)
	duc := controllers.NewDeleteUserController(*du)

	user := r.Group("users")
	{
		user.POST("/", ruc.Execute)
		user.GET("/:id", gouc.Execute)
		user.PUT("/user/:id", euc.Execute)
		user.PUT("/credential/:id", ecc.Execute)
		user.DELETE("/:id", duc.Execute)
	}
}