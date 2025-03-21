package infrastructure

import (
	"log"

	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/core"
)

type ParamsMySQL struct {
	conn core.ConectionMySQL
}

func NewParamsMySQL()*ParamsMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&ParamsMySQL{conn: *conn}
}

func(r *ParamsMySQL)SetParams(params models.CultivationParameters)(int, error){
	query := "INSERT INTO cultivation_parameters (humidity_min, humidity_max, temp_min, temp_max, min_air_con, max_air_con) VALUES (?,?,?,?,?,?)"
	res, err := r.conn.ExecPreparedQuerys(
		query, 
		params.Humidity_min, 
		params.Humidity_max, 
		params.Temp_min, 
		params.Temp_max, 
		params.Min_air_con, 
		params.Max_air_con,
	)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	last_id_param, err := res.LastInsertId()
	if err != nil {
        log.Fatalf("Error al obtener último usuario registrado:", err)
    }
	return int(last_id_param), err
}

func(r *ParamsMySQL)GetParams(id int)([]models.CultivationParameters, error){
	query := "SELECT * FROM cultivation_parameters WHERE id_cultivation_parameter = ?"
	rows, err := r.conn.FetchRows(query, id)
	var params []models.CultivationParameters
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var min_hum float32
		var max_hum float32
		var min_temp float32
		var max_temp float32
		var min_air float32
		var max_air float32

		if err := rows.Scan(&id, &min_hum, &max_hum, &min_temp, &max_temp, &min_air, &max_air); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		param := &models.CultivationParameters{
			ID: id, 
			Humidity_min: min_hum, 
			Humidity_max: max_hum,
			Temp_min: min_temp,
			Temp_max: max_temp,
			Min_air_con: min_air,
			Max_air_con: max_air,
		}
		params = append(params, *param)
	}
	return params, err
}

func(r *ParamsMySQL)EditParameters(params models.CultivationParameters)error{
	query := "UPDATE cultivation_parameters SET humidity_min = ?, humidity_max = ?, temp_min = ?, temp_max = ?, min_air_con = ?,  max_air_con = ? WHERE id_cultivation_parameter = ?"
	_, err := r.conn.ExecPreparedQuerys(
		query, 
		params.Humidity_min, 
		params.Humidity_max, 
		params.Temp_min, 
		params.Temp_max, 
		params.Min_air_con, 
		params.Max_air_con,
		params.ID,
	)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err	
}

func(r *ParamsMySQL)Deleteparams(id int)error{
	query := "DELETE FROM cultivation_parameters WHERE id_cultivation_parameter = ?"
	_, err := r.conn.ExecPreparedQuerys(query, id)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err	
}