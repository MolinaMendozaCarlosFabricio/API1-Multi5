package models

type ParcelAllInfo struct{
	ID int
	Id_user int
	Id_crop CropAllInfo
	Id_device int
	Id_status CropStatus
}