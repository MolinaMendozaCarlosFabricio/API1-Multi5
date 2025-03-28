package infrastructure

var measurement_mysql MeasurementMySQL

func MeasurementsDependences(){
	measurement_mysql = *NewMeasurementMySQL()
}

func GetMeasurementMySQL()*MeasurementMySQL{
	return &measurement_mysql
}