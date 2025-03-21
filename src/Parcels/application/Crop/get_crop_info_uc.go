package crop

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type GetCropInfoUC struct {
	db repository.ICrop_repo
}

func NewGetCropInfoUC(db repository.ICrop_repo)*GetCropInfoUC{
	return&GetCropInfoUC{db: db}
}

func(uc *GetCropInfoUC)Execute(id int)([]models.Crop, error){
	return uc.db.GetCropInfo(id)
}