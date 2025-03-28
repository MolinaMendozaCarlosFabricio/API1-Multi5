package devicehistory

import (
	"time"

	"api1-multi.com/a/src/Devices/domain/repository"
)

type SetRetireDateUC struct {
	db repository.IDeviceHistory_repo
}

func NewSetRetireDateUC(db repository.IDeviceHistory_repo)*SetRetireDateUC{
	return&SetRetireDateUC{db: db}
}

func(uc *SetRetireDateUC)Execute(id int)error{
	date := time.Now().Format("2006-01-02")
	return uc.db.SetRetireDate(id, date)
}