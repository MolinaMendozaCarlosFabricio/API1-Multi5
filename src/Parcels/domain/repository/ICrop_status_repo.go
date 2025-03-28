package repository

import "api1-multi.com/a/src/Parcels/domain/models"

type ICropSatusRepo interface {
	GetCropStatus(id int)([]models.CropStatus, error)
}