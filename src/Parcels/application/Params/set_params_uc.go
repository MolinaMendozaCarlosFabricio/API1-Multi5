package params

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type SetParamsUC struct {
	db repository.IParameters_repo
}

func NewSetParamsUC(db repository.IParameters_repo)*SetParamsUC{
	return&SetParamsUC{db: db}
}

func(uc *SetParamsUC)Execute(
	humidity_min, 
	humidity_max, 
	temp_min, 
	temp_max, 
	min_air_con, 
	max_air_con float32,
)(int, error){
	param := &models.CultivationParameters{
		ID: 0, 
		Humidity_min: humidity_min,
		Humidity_max: humidity_max,
		Temp_min: temp_min,
		Temp_max: temp_max,
		Min_air_con: min_air_con,
		Max_air_con: max_air_con,
	}
	return uc.db.SetParams(*param)
}