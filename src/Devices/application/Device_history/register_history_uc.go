package devicehistory

import (
	"time"

	"api1-multi.com/a/src/Devices/domain/models"
	"api1-multi.com/a/src/Devices/domain/repository"
)

type RegisterDeviceInHistoryUC struct {
	db repository.IDeviceHistory_repo
}

func NewRegisterDeviceInHistoryUC(db repository.IDeviceHistory_repo)*RegisterDeviceInHistoryUC{
	return&RegisterDeviceInHistoryUC{db: db}
}

func(uc *RegisterDeviceInHistoryUC)Execute(id_parcel int)error{
	history_record := models.DeviceHistory{
		ID: 0, 
		Id_parcel: id_parcel, 
		Date_installation: time.Now().Format("2006-01-02"),
		Date_maintenance: "",
		Date_retire: "",
		Device_status: "",
	}
	return uc.db.RegisterDeviceInHistory(history_record)
}