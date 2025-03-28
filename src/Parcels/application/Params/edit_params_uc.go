package params

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type EditParametersUc struct {
	db repository.IParameters_repo
}

func NewEditParametersUc(db repository.IParameters_repo)*EditParametersUc{
	return&EditParametersUc{db: db}
}

func(uc *EditParametersUc)Execute(
	id int,
	humidity_min, 
	humidity_max, 
	temp_min, 
	temp_max, 
	min_air_con, 
	max_air_con float32,
)error{
	param := &models.CultivationParameters{
		ID: id, 
		Humidity_min: humidity_min,
		Humidity_max: humidity_max,
		Temp_min: temp_min,
		Temp_max: temp_max,
		Min_air_con: min_air_con,
		Max_air_con: max_air_con,
	}
	return uc.db.EditParameters(*param)
}