package application

import (
	"time"

	"api1-multi.com/a/src/Measurement/domain"
)

type RegisterMeasurementUC struct {
	db domain.IMeasurement_repo
}

func NewRegisterMeasurementUC(db domain.IMeasurement_repo)*RegisterMeasurementUC{
	return&RegisterMeasurementUC{db: db}
}

func(uc *RegisterMeasurementUC)Execute(id_parcel int, temp, humedity, air, sun float32)error{
	measurement := &domain.Measurement{
		ID: 0, 
		Id_parcel: id_parcel, 
		Temp: temp, 
		Humedity: humedity, 
		Air: air, 
		Sun: sun,
		Date_and_hour: time.Now().Format("2006-01-02 15:04:05"),
	}
	return uc.db.RegisterMeasurement(*measurement)
}