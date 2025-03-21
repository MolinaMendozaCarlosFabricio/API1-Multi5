package repository

import "api1-multi.com/a/src/Parcels/domain/models"

type ICropType_repo interface {
	GetCropType(id int)([]models.CropType, error)
}