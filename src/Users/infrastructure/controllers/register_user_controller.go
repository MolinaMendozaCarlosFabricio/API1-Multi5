package controllers

import (
	"net/http"

	"api1-multi.com/a/src/Users/application"
	"api1-multi.com/a/src/Users/infrastructure"
	"github.com/gin-gonic/gin"
)

type RegisterUserController struct {
	c application.RegisterUserUC
}

func NewRegisterUserController()*RegisterUserController{

	mysql := infrastructure.GetUserMySQL()
	uc := application.NewRegisterUserUC(mysql)

	return&RegisterUserController{c: *uc}
}

func(controller *RegisterUserController)Execute(c *gin.Context){
	var input struct {
		First_name string `json:"first_name"`
		Last_name string `json:"last_name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.c.Execute(
		input.First_name, input.Last_name, input.Email, input.Password,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Usuario registrado",
	})
}