package repository

import "api1-multi.com/a/src/Devices/domain/models"

type IDeviceHistory_repo interface {
	RegisterDeviceInHistory(history models.DeviceHistory)error
	GetDeviceHistory(id_parcel int)([]models.DeviceHistory, error)
	SetMaintenanceDate(id int, date string)error
	SetRetireDate(id int, date string)error
}