package repository

import "api1-multi.com/a/src/Devices/domain/models"

type IDevice_repo interface {
	GetDevice(id int)([]models.Device, error)
}