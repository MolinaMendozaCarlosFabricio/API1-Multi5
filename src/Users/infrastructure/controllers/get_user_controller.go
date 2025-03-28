package controllers

import (
	"net/http"
	"strconv"

	"api1-multi.com/a/src/Users/application"
	"api1-multi.com/a/src/Users/infrastructure"
	"github.com/gin-gonic/gin"
)

type ViewUserController struct {
	c application.ViewUserUC
}

func NewViewUserController()*ViewUserController{

	mysql := infrastructure.GetUserMySQL()
	uc := application.NewViewUserUC(mysql)

	return&ViewUserController{c: *uc}
}

func(controller *ViewUserController)Execute(c *gin.Context){
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

	results, err := controller.c.Execute(id_number)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Usuarios obtenidos",
		"Results": results,
	})
}