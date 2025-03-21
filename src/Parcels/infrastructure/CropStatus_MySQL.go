package infrastructure

import (
	"log"

	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/core"
)

type CropStatusMySQL struct {
	conn core.ConectionMySQL
}

func NewCropStatusMySQL()*CropStatusMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&CropStatusMySQL{conn: *conn}
}

func(r *CropStatusMySQL)GetCropStatus(id int)([]models.CropStatus, error){
	query := "SELECT * FROM crop_status WHERE id_crop_status = ?"
	rows, err := r.conn.FetchRows(query, id)
	var crop_status []models.CropStatus
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var crop_statu models.CropStatus
		
		if err := rows.Scan(&crop_statu.ID, &crop_statu.Name); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		
		crop_status = append(crop_status, crop_statu)
	}
	return crop_status, err
}