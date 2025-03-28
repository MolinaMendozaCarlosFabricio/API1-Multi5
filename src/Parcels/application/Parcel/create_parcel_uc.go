package parcel

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type CreateParcelUC struct {
	db repository.IParcel_repo
}

func NewCreateParcelUC(db repository.IParcel_repo)*CreateParcelUC{
	return&CreateParcelUC{db: db}
}

func(uc *CreateParcelUC)Execute(id_user, id_crop, id_device, id_status int)(int, error){
	parcel := &models.Parcel{ID: 0, Id_user: id_user, Id_crop: id_crop, Id_device: id_device, Id_status: id_status}
	return uc.db.CreateParcel(*parcel)
}