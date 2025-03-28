package devicehistory

import (
	"time"

	"api1-multi.com/a/src/Devices/domain/repository"
)

type SetMaintenanceDateUC struct {
	db repository.IDeviceHistory_repo
}

func NewSetMaintenanceDateUC(db repository.IDeviceHistory_repo)*SetMaintenanceDateUC{
	return&SetMaintenanceDateUC{db: db}
}

func(uc *SetMaintenanceDateUC)Execute(id int)error{
	date := time.Now().Format("2006-01-02")
	return uc.db.SetMaintenanceDate(id, date)
}