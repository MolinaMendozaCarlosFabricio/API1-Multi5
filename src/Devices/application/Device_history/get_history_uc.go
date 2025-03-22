package devicehistory

import (
	"api1-multi.com/a/src/Devices/domain/models"
	"api1-multi.com/a/src/Devices/domain/repository"
)

type GetDeviceHistoryUC struct {
	db repository.IDeviceHistory_repo
}

func NewGetDeviceHistoryUC(db repository.IDeviceHistory_repo)*GetDeviceHistoryUC{
	return&GetDeviceHistoryUC{db: db}
}

func(uc *GetDeviceHistoryUC)Execute(id_parcel int)([]models.DeviceHistory, error){
	return uc.db.GetDeviceHistory(id_parcel)
}