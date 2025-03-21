package params

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type GetParamsUC struct {
	db repository.IParameters_repo
}

func NewGetParamsUC(db repository.IParameters_repo)*GetParamsUC{
	return&GetParamsUC{db: db}
}

func(uc *GetParamsUC)Execute(id int)([]models.CultivationParameters, error){
	return uc.db.GetParams(id)
}