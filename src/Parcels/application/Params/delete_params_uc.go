package params

import "api1-multi.com/a/src/Parcels/domain/repository"

type DeleteParamsUC struct {
	db repository.IParameters_repo
}

func NewDeleteParamsUC(db repository.IParameters_repo)*DeleteParamsUC{
	return&DeleteParamsUC{db: db}
}

func(uc *DeleteParamsUC)Execute(id int)error{
	return uc.db.Deleteparams(id)
}