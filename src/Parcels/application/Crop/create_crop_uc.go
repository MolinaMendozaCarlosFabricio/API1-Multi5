package crop

import (
	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/Parcels/domain/repository"
)

type CreateCropUC struct {
	db repository.ICrop_repo
}

func NewCreateCropUC(db repository.ICrop_repo)*CreateCropUC{
	return&CreateCropUC{db: db}
}

func(uc *CreateCropUC)Execute( 
	name string, 
	id_cultivation_parameter int,
	Id_crop_type int,
)(int, error){
	crop := &models.Crop{
		ID: 0, 
		Name: name, 
		Id_cultivation_parameter: id_cultivation_parameter, 
		Id_crop_type: Id_crop_type,
	}
	return uc.db.CreateCrop(*crop)
}