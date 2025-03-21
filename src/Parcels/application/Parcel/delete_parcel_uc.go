package parcel

import "api1-multi.com/a/src/Parcels/domain/repository"

type DeleteParcelUC struct {
	db repository.IParcel_repo
}

func NewDeleteParcelUC(db repository.IParcel_repo)*DeleteParcelUC{
	return&DeleteParcelUC{db: db}
}

func(uc *DeleteParcelUC)Execute(id int)error{
	return uc.db.DeleteParcel(id)
}