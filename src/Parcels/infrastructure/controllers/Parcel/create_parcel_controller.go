package parcel

import (
	"net/http"

	crop "api1-multi.com/a/src/Parcels/application/Crop"
	params "api1-multi.com/a/src/Parcels/application/Params"
	parcel "api1-multi.com/a/src/Parcels/application/Parcel"
	"api1-multi.com/a/src/Parcels/infrastructure"
	"github.com/gin-gonic/gin"
)

type CreateParcelC struct {
	uc_crop crop.CreateCropUC
	uc_params params.SetParamsUC
	uc_parcel parcel.CreateParcelUC
}

func NewCreateParcelC()*CreateParcelC{

	crop_mysql := infrastructure.GetCropMySQL()
	params_mysql := infrastructure.GetParamsMySQL()
	parcel_mysql := infrastructure.GetParcelMysQL()

	uc_crop := crop.NewCreateCropUC(crop_mysql)
	uc_params := params.NewSetParamsUC(params_mysql)
	uc_parcel := parcel.NewCreateParcelUC(parcel_mysql)

	return&CreateParcelC{uc_crop: *uc_crop, uc_params: *uc_params, uc_parcel: *uc_parcel}
}

func(controller *CreateParcelC)Execute(c *gin.Context){
	var input struct{
		Humidity_min float32 `json:"humidity_min"`
		Humidity_max float32 `json:"humidity_max"`
		Temp_min float32 `json:"temp_min"`
		Temp_max float32 `json:"temp_max"`
		Min_air_con float32 `json:"min_air_con"`
		Max_air_con float32 `json:"max_air_con"`
		Name string `json:"name"`
		Id_crop_type int `json:"id_crop_type"`
		Id_user int `json:"id_user"`
		Id_device int `json:"id_device"`
		Id_status int `json:"id_status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	id_parameter, err := controller.uc_params.Execute(
		input.Humidity_min,
		input.Humidity_max,
		input.Temp_min,
		input.Temp_max,
		input.Min_air_con,
		input.Max_air_con,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al ingresar parámetros de cultivo",
		})
		return
	}

	id_crop, err := controller.uc_crop.Execute(input.Name, id_parameter, input.Id_crop_type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar información del cultivo",
		})
		return
	}

	results, err := controller.uc_parcel.Execute(input.Id_user, id_crop, input.Id_device, input.Id_status);

	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar parcela",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Parcela registrada",
		"Results": results,
	})
}