package models

type CropAllInfo struct{
	ID int
	Name string
	Id_cultivation_parameter CultivationParameters
	Id_crop_type CropType
}