package infrastructure

import (
	"log"

	"api1-multi.com/a/src/Devices/domain/models"
	"api1-multi.com/a/src/core"
)

type DeviceHistoryMySQL struct {
	conn core.ConectionMySQL
}

func NewDeviceHistoryMySQL()*DeviceHistoryMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&DeviceHistoryMySQL{conn: *conn}
}

func(r *DeviceHistoryMySQL)RegisterDeviceInHistory(history models.DeviceHistory)error{
	query := "INSERT INTO device_history (id_parcel, date_instalation, id_device_status) VALUES (?, ?, 1)"
	_, err := r.conn.ExecPreparedQuerys(
		query,
		history.Id_parcel,
		history.Date_installation,
	)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}

func(r *DeviceHistoryMySQL)GetDeviceHistory(id_parcel int)([]models.DeviceHistory, error){
	query := "SELECT device_history.id_device_history, device_history.date_instalation, device_history.date_mainentance, device_history.date_retire, device_status.name_status FROM device_history INNER JOIN device_status ON device_history.id_device_status = device_status.id_device_status WHERE device_history.id_parcel = ?"
	rows, err := r.conn.FetchRows(query, id_parcel)
	var history []models.DeviceHistory
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var date_installation string
		var date_maintenance string
		var date_retire string
		var status string
		if err := rows.Scan(&id, &date_installation, &date_maintenance, &date_retire, &status); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		device := &models.DeviceHistory{
			ID: id, 
			Id_parcel: id_parcel,
			Date_installation: date_installation,
			Date_maintenance: date_maintenance,
			Date_retire: date_retire,
			Device_status: status,
		}
		history = append(history, *device)
	}
	return history, err
}

func(r *DeviceHistoryMySQL)SetMaintenanceDate(id int, date string)error{
	query := "UPDATE device_history SET date_mainentance = ? WHERE id_device_history = ?"
	_, err := r.conn.ExecPreparedQuerys(query, date, id)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}

func(r *DeviceHistoryMySQL)SetRetireDate(id int, date string)error{
	query := "UPDATE device_history SET date_retire = ?, id_device_status = 2 WHERE id_device_history = ?"
	_, err := r.conn.ExecPreparedQuerys(query, date, id)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}