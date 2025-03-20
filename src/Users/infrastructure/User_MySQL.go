package infrastructure

import (
	"log"

	"api1-multi.com/a/src/core"
	"api1-multi.com/a/src/Users/domain"
)

type UserMySQL struct {
	conn core.ConectionMySQL
}

func NewUserMySQL()*UserMySQL{
	conn := core.MySQLConection()
    if conn.Err != "" {
        log.Fatal("Error al configurar la pool de conexiones", conn.Err)
    }
	return&UserMySQL{conn: *conn}
}

func(r *UserMySQL)RegisterUserMethod(user domain.User, credentials domain.AccessCredentials)error{
	query := "INSERT INTO users (first_name, last_name) VALUES (?, ?)"
	res, err := r.conn.ExecPreparedQuerys(query, user.First_name, user.Last_name)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
	last_id, err := res.LastInsertId()
	if err != nil {
        log.Fatalf("Error al obtener último usuario registrado:", err)
    }
	query = "INSERT INTO access_credentials (id_user, email, password) VALUES (?, ?, ?)"
	res, err = r.conn.ExecPreparedQuerys(query, last_id, credentials.Email, credentials.Password)
	if err != nil {
        log.Fatalf("Error al registrar Usuarios:", err)
    }
    return err
}

func(r *UserMySQL)ViewUserMethod(id int)([]domain.User, error){
	query := "SELECT * FROM users WHERE id_user = ?"
	rows, err := r.conn.FetchRows(query, id)
	var users []domain.User
	if err != nil {
        log.Fatalf("Error al obtener Usuarios:", err)
    }
    defer rows.Close()
	for rows.Next(){
		var id int
		var first_name string
		var last_name string

		if err := rows.Scan(&id, &first_name, &last_name); err != nil{
			log.Println("Error al escanear la fila:", err)
		}
		user := domain.User{ID: id, First_name: first_name, Last_name: last_name}
		users = append(users, user)
	}
	return users, err
}

func(r *UserMySQL)EditUserMethod(user domain.User)error{
	query := "UPDATE users SET first_name = ?, last_name = ? WHERE id_user = ?"
	_, err := r.conn.ExecPreparedQuerys(query, user.First_name, user.Last_name, user.ID)
	if err != nil {
        log.Fatalf("Error al editar usuario:", err)
    }
    return err
}

func(r *UserMySQL)EditCredentialsMethod(credentials domain.AccessCredentials)error{
	query := "UPDATE access_credentials SET email = ?, password = ? WHERE id_user = ?"
	_, err := r.conn.ExecPreparedQuerys(query, credentials.Email, credentials.Password, credentials.Id_user)
	if err != nil {
        log.Fatalf("Error al editar credenciales:", err)
    }
    return err
}

func(r *UserMySQL)DeleteUserMethod(id int)error{
	query := "DELETE FROM users WHERE id_user = ?"
	_, err := r.conn.ExecPreparedQuerys(query, id)
	if err != nil {
        log.Fatalf("Error al eliminar usuario:", err)
    }
	return err
}