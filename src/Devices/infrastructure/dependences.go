package infrastructure

var device_mysql DeviceMySQL
var device_history_mysql DeviceHistoryMySQL

func DevicesDependences(){
	device_mysql = *NewDeviceMySQL()
	device_history_mysql = *NewDeviceHistoryMySQL()
}

func GetDeviceMySQL()*DeviceMySQL{
	return &device_mysql
}

func GetDeviceHistoryMySQL()*DeviceHistoryMySQL{
	return &device_history_mysql
}