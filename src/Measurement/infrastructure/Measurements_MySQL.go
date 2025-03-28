package infrastructure

import (
	"log"

	"api1-multi.com/a/src/Measurement/domain"
	"api1-multi.com/a/src/core"
)

type MeasurementMySQL struct {
	conn core.ConectionMySQL
}

func NewMeasurementMySQL()*MeasurementMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&MeasurementMySQL{conn: *conn}
}

func(r *MeasurementMySQL)RegisterMeasurement(measurement domain.Measurement)error{
	query := "INSERT INTO measurements (id_parcel, temp, humedity, air, sun, date_and_hour) VALUES (?,?,?,?,?,?)"
	_, err := r.conn.ExecPreparedQuerys(
		query,
		measurement.Id_parcel,
		measurement.Temp,
		measurement.Humedity,
		measurement.Air,
		measurement.Sun,
		measurement.Date_and_hour,
	)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}

func(r *MeasurementMySQL)GetAllMeasurements(id_parcel int)([]domain.Measurement, error){
	query := "SELECT * FROM measurements WHERE id_parcel = ?"
	rows, err := r.conn.FetchRows(query, id_parcel)
	var measurements []domain.Measurement
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var id_parcel int
		var temp float32
		var humedity float32
		var air float32
		var sun float32
		
		if err := rows.Scan(&id, &id_parcel, &temp, &humedity, &air, &sun); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		measurement := &domain.Measurement{
			ID: id,
			Id_parcel: id_parcel,
			Temp: temp,
			Humedity: humedity,
			Air: air,
			Sun: sun,
		}
		measurements = append(measurements, *measurement)
	}
	return measurements, err
}

func(r *MeasurementMySQL)GetOneMeasurement(id int)([]domain.Measurement, error){
	query := "SELECT * FROM measurements WHERE id_measurement = ?"
	rows, err := r.conn.FetchRows(query, id)
	var measurements []domain.Measurement
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var id_parcel int
		var temp float32
		var humedity float32
		var air float32
		var sun float32
		
		if err := rows.Scan(&id, &id_parcel, &temp, &humedity, &air, &sun); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		measurement := &domain.Measurement{
			ID: id,
			Id_parcel: id_parcel,
			Temp: temp,
			Humedity: humedity,
			Air: air,
			Sun: sun,
		}
		measurements = append(measurements, *measurement)
	}
	return measurements, err
}