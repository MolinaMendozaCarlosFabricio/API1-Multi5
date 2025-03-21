package cropstatus

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type GetCropStatusUC struct {
	db repository.ICropSatusRepo
}

func NewGetCropStatusUC(db repository.ICropSatusRepo)*GetCropStatusUC{
	return&GetCropStatusUC{db: db}
}

func(uc *GetCropStatusUC)Execute(id int)([]models.CropStatus, error){
	return uc.db.GetCropStatus(id)
}