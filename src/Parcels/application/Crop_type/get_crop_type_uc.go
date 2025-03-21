package croptype

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type GetCropTypeUC struct {
	db repository.ICropType_repo
}

func NewGetCropTypeUC(db repository.ICropType_repo)*GetCropTypeUC{
	return&GetCropTypeUC{db: db}
}

func(uc *GetCropTypeUC)Execute(id int)([]models.CropType, error){
	return uc.db.GetCropType(id)
}