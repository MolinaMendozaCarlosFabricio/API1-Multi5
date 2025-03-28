package parcel

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type GetAllMyParcelsUC struct {
	db repository.IParcel_repo
}

func NewGetAllMyParcelsUC(db repository.IParcel_repo)*GetAllMyParcelsUC{
	return&GetAllMyParcelsUC{db: db}
}

func(uc *GetAllMyParcelsUC)Execute(id_user int)([]models.ParcelAllInfo, error){
	return uc.db.GetAllMyParcels(id_user)
}