package models

type CultivationParameters struct{
	ID int
	Humidity_min float32
	Humidity_max float32
	Temp_min float32
	Temp_max float32
	Min_air_con float32
	Max_air_con float32
}