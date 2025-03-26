package infrastructure

var mysql UserMySQL

func UserDependences(){
	mysql = *NewUserMySQL()
}

func GetUserMySQL()*UserMySQL{
	return &mysql
}