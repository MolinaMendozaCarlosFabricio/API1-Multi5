package application

import "api1-multi.com/a/src/Measurement/domain"

type GetAllMeasurementsUC struct {
	db domain.IMeasurement_repo
}

func NewGetAllMeasurementsUC(db domain.IMeasurement_repo)*GetAllMeasurementsUC{
	return&GetAllMeasurementsUC{db: db}
}

func(uc *GetAllMeasurementsUC)Execute(id_parcel int)([]domain.Measurement, error){
	return uc.db.GetAllMeasurements(id_parcel)
}