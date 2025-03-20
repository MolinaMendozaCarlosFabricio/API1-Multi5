package infrastructure

import (
	"log"

	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/core"
)

type CropMySQL struct {
	conn core.ConectionMySQL
}

func NewCropMySQL()*CropMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&CropMySQL{conn: *conn}
}

func(r *CropMySQL)CreateCrop(crop models.Crop)(int, error){
	query := "INSERT INTO crop (crop_name, id_cultivation_parameter, id_crop_type) VALUES (?,?,?)"
	res, err := r.conn.ExecPreparedQuerys(query, crop.Name, crop.Id_cultivation_parameter, crop.Id_crop_type)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	last_id, err := res.LastInsertId()
	if err != nil {
        log.Fatalf("Error al obtener último usuario registrado:", err)
    }
	return int(last_id), err
}

func(r *CropMySQL)EditNameCrop(id int, new_name string)error{
	query := "INSERT INTO crop (crop_name, id_cultivation_parameter, id_crop_type) VALUES (?,?,?)"
	_, err := r.conn.ExecPreparedQuerys(query)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}