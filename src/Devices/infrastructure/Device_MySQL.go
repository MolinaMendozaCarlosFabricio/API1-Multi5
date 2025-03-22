package infrastructure

import (
	"log"

	"api1-multi.com/a/src/Devices/domain/models"
	"api1-multi.com/a/src/core"
)

type DeviceMySQL struct {
	conn core.ConectionMySQL
}

func NewDeviceMySQL()*DeviceMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&DeviceMySQL{conn: *conn}
}

func(r *DeviceMySQL)GetDevice(id int)([]models.Device, error){
	query := "SELECT * FROM devices WHERE id_device = ?"
	rows, err := r.conn.FetchRows(query, id)
	var devices []models.Device
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var model string
		
		if err := rows.Scan(&id, &model); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		device := &models.Device{ID: id, Model: model}
		devices = append(devices, *device)
	}
	return devices, err
}