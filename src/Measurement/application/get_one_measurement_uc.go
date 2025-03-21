package application

import "api1-multi.com/a/src/Measurement/domain"

type GetOneMeasurementUC struct {
	db domain.IMeasurement_repo
}

func NewGetOneMeasurementUC(db domain.IMeasurement_repo)*GetOneMeasurementUC{
	return&GetOneMeasurementUC{db: db}
}

func(uc *GetOneMeasurementUC)Execute(id int)([]domain.Measurement, error){
	return uc.db.GetOneMeasurement(id)
}