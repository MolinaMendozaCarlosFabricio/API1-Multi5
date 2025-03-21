package repository

import "api1-multi.com/a/src/Parcels/domain/models"

type IParameters_repo interface {
	SetParams(params models.CultivationParameters)(int, error)
	GetParams(id int)([]models.CultivationParameters, error)
	EditParameters(params models.CultivationParameters)error
	Deleteparams(id int)error
}