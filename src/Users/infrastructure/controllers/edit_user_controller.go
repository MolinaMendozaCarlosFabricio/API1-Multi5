package controllers

import (
	"net/http"
	"strconv"

	"api1-multi.com/a/src/Users/application"
	"api1-multi.com/a/src/Users/infrastructure"
	"github.com/gin-gonic/gin"
)

type EditUserController struct {
	c application.EditUserUC
}

func NewEditUserController()*EditUserController{

	mysql := infrastructure.GetUserMySQL()
	uc := application.NewEditUserUC(mysql)

	return&EditUserController{c: *uc}
}

func(controller *EditUserController)Execute(c *gin.Context){
	id, error_param := c.Params.Get("id")
	if !error_param {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "No se pudo mapear el parámetro",
		})
		return
	}

	id_number, error_strconv := strconv.Atoi(id)
	if error_strconv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Parámetro incorrecto",
		})
		return
	}

	var input struct {
		First_name string `json:"first_name"`
		Last_name string `json:"last_name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.c.Execute(id_number, input.First_name, input.Last_name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al editar usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Usuario editado",
	})
}