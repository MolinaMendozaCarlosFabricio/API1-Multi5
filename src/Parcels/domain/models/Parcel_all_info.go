package models

import "api1-multi.com/a/src/Devices/domain/models"

type ParcelAllInfo struct {
	ID        int
	Id_user   int
	Id_crop   CropAllInfo
	Id_device models.Device
	Id_status CropStatus
}