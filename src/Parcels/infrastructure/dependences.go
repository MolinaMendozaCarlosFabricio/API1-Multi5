package infrastructure

var crop_mysql CropMySQL
var crop_status_mysql CropStatusMySQL
var crop_type_mysql CropTypeMySQL
var params_mysql ParamsMySQL
var parcel_mysql ParcelMySQL

func ParcelsDependences(){
	crop_mysql = *NewCropMySQL()
	crop_status_mysql = *NewCropStatusMySQL()
	crop_type_mysql = *NewCropTypeMySQL()
	params_mysql = *NewParamsMySQL()
	parcel_mysql = *NewParcelMySQL()
}

func GetCropMySQL()*CropMySQL{
	return &crop_mysql
}

func GetCropStatusMySQL()*CropStatusMySQL{
	return &crop_status_mysql
}

func GetCropSTypeMySQL()*CropTypeMySQL{
	return &crop_type_mysql
}

func GetParamsMySQL()*ParamsMySQL{
	return &params_mysql
}

func GetParcelMysQL()*ParcelMySQL{
	return &parcel_mysql
}