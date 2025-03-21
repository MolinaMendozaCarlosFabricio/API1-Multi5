package infrastructure

import (
	"log"

	"api1-multi.com/a/src/Parcels/domain/models"
	"api1-multi.com/a/src/core"
)

type ParcelMySQL struct {
	conn core.ConectionMySQL
}

func NewParcelMySQL()*ParcelMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&ParcelMySQL{conn: *conn}
}

func(r *ParcelMySQL)CreateParcel(parcel models.Parcel)error{
	query := "INSERT INTO parcel (id_user, id_crop, id_device, id_status) VALUES (?,?,?,?)"
	_, err := r.conn.ExecPreparedQuerys(query, parcel.Id_user, parcel.Id_crop, parcel.Id_device, parcel.Id_status)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}

func(r *ParcelMySQL)GetAllMyParcels(id_user int)([]models.ParcelAllInfo, error){
	query := "SELECT parcel.id_parcel, crop.crop_name, crop_type.crop_type_name, crop_status.name, devices.model FROM parcel INNER JOIN crop ON parcel.id_crop = crop.id_crop INNER JOIN crop_type ON crop_type.id_crop_type = crop.id_crop_type INNER JOIN crop_status ON crop_status.id_crop_status = parcel.id_status INNER JOIN devices ON devices.id_device = parcel.id_device WHERE id_user = ?;"
	rows, err := r.conn.FetchRows(query, id_user)
	var parcels []models.ParcelAllInfo
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var name string
		var type_name string
		var status string
		var model string
		
		if err := rows.Scan(&id, &name, &type_name, &status, &model); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		var parcel models.ParcelAllInfo
		parcel.ID = id
		parcel.Id_crop.Name = name
		parcel.Id_crop.Id_crop_type.Name = type_name
		parcel.Id_status.Name = status
		parcel.Id_device.Model = model
		parcels = append(parcels, parcel)
	}
	return parcels, err
}

func(r *ParcelMySQL)GetAllAboutMyParcel(id int)([]models.Parcel, error){
	query := "SELECT * FROM parcel WHERE id_parcel = ?"
	rows, err := r.conn.FetchRows(query, id)
	var parcels []models.Parcel
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var parcel models.Parcel
		
		if err := rows.Scan(&parcel.ID, &parcel.Id_user, &parcel.Id_crop, &parcel.Id_device, &parcel.Id_status); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		
		parcels = append(parcels, parcel)
	}
	return parcels, err
}

func(r *ParcelMySQL)EditParcel(parcel models.Parcel)error{
	query := "UPDATE parcel SET id_device = ?, id_status = ? WHERE id_parcel = ?"
	_, err := r.conn.ExecPreparedQuerys(query, parcel.Id_device, parcel.Id_status, parcel.ID)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}

func(r *ParcelMySQL)DeleteParcel(id int)error{
	query := "DELETE FROM parcel WHERE id_parcel = ?"
	_, err := r.conn.ExecPreparedQuerys(query, id)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	return err
}