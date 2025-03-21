package device

import (
	"api1-multi.com/a/src/Devices/domain/models"
	"api1-multi.com/a/src/Devices/domain/repository"
)

type GetDeviceUC struct {
	db repository.IDevice_repo
}

func NewGetDeviceUC(db repository.IDevice_repo)*GetDeviceUC{
	return&GetDeviceUC{db: db}
}

func(uc *GetDeviceUC)Execute(id int)([]models.Device, error){
	return uc.db.GetDevice(id)
}