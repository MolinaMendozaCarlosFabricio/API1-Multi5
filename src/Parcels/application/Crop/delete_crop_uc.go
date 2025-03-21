package crop

import "api1-multi.com/a/src/Parcels/domain/repository"

type DeleteCropUC struct {
	db repository.ICrop_repo
}

func NewDeleteCropUC(db repository.ICrop_repo)*DeleteCropUC{
	return&DeleteCropUC{db: db}
}

func(uc *DeleteCropUC)Execute(id int)error{
	return uc.db.DeleteCrop(id)
}