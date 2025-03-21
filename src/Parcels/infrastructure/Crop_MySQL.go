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

func(r *CropMySQL)GetCropInfo(id int)([]models.Crop, error){
	query := "SELECT * FROM parcel WHERE id_parcel = ?"
	rows, err := r.conn.FetchRows(query, id)
	var crops []models.Crop
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var crop models.Crop
		
		if err := rows.Scan(&crop.ID, &crop.Name, &crop.Id_cultivation_parameter, &crop.Id_crop_type); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		
		crops = append(crops, crop)
	}
	return crops, err
}

func(r *CropMySQL)EditNameCrop(id int, new_name string)error{
	query := "UPDATE crop SET crop_name = ? WHERE id_crop = ?"
	_, err := r.conn.ExecPreparedQuerys(query, new_name, id)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}

func(r *CropMySQL)DeleteCrop(id int)error{
	query := "DELETE FROM crop WHERE id_crop = ?"
	_, err := r.conn.ExecPreparedQuerys(query, id)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}