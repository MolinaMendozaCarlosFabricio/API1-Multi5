package parcel

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type EditParcelUC struct {
	db repository.IParcel_repo
}

func NewEditParcelUC(db repository.IParcel_repo)*EditParcelUC{
	return&EditParcelUC{db: db}
}

func(uc *EditParcelUC)Execute(id, id_user, id_crop, id_device, id_status int)error{
	parcel := &models.Parcel{
		ID: id, 
		Id_user: id_user, 
		Id_crop: id_crop, 
		Id_device: id_device, 
		Id_status: id_status,
	}
	return uc.db.EditParcel(*parcel)
}