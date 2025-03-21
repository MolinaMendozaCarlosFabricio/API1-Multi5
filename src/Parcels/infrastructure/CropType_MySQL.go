package infrastructure

import (
	"log"

	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/core"
)

type CropTypeMySQL struct {
	conn core.ConectionMySQL
}

func NewCropTypeMySQL()*CropTypeMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&CropTypeMySQL{conn: *conn}
}

func(r *CropTypeMySQL)GetCropType(id int)([]models.CropType, error){
	query := "SELECT * FROM crop_type WHERE id_crop_type = ?"
	rows, err := r.conn.FetchRows(query, id)
	var crop_types []models.CropType
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var crop_type models.CropType
		
		if err := rows.Scan(&crop_type.ID, &crop_type.Name, &crop_type.Description); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		
		crop_types = append(crop_types, crop_type)
	}
	return crop_types, err
}