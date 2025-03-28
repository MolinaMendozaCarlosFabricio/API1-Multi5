package domain

type IMeasurement_repo interface{
	RegisterMeasurement(measurement Measurement)error
	GetAllMeasurements(id_parcel int)([]Measurement, error)
	GetOneMeasurement(id int)([]Measurement, error)
}