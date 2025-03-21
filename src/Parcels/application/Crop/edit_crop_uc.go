package crop

import "api1-multi.com/a/src/Parcels/domain/repository"

type EditNameCropUC struct {
	db repository.ICrop_repo
}

func NewEditNameCropUC(db repository.ICrop_repo)*EditNameCropUC{
	return&EditNameCropUC{db: db}
}

func(uc *EditNameCropUC)Execute(id int, name string)error{
	return uc.db.EditNameCrop(id, name)
}