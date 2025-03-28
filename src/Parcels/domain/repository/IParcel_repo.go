package repository

import "api1-multi.com/a/src/Parcels/domain/models"

type IParcel_repo interface {
	CreateParcel(parcel models.Parcel)(int, error)
	GetAllMyParcels(id_user int)([]models.ParcelAllInfo, error)
	GetAllAboutMyParcel(id int)([]models.Parcel, error)
	EditParcel(parcel models.Parcel)error
	DeleteParcel(id int)error
}