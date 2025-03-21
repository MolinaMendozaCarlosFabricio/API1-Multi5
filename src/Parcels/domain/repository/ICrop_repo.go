package repository

import "api1-multi.com/a/src/Parcels/domain/models"

type ICrop_repo interface {
	CreateCrop(crop models.Crop)(int, error)
	GetCropInfo(id int)([]models.Crop, error)
	EditNameCrop(id int, new_name string)error
	DeleteCrop(id int)error
}