package parcel

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type GetAllAboutMyParcelUC struct {
	db repository.IParcel_repo
}

func NewGetAllAboutMyParcel(db repository.IParcel_repo)*GetAllAboutMyParcelUC{
	return&GetAllAboutMyParcelUC{db: db}
}

func(uc *GetAllAboutMyParcelUC)Execute(id int)([]models.Parcel, error){
	return uc.db.GetAllAboutMyParcel(id)
}