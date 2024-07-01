package models

import (
	"Tubes_PBW/config"
	"Tubes_PBW/entities"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SewaModel struct {
    conn *sql.DB
}

func NewSewaModel() *SewaModel {
    conn, err := config.DBconnection()
    if err != nil {
        panic(err)
    }

    return &SewaModel{
        conn: conn,
    }
}

func (s *SewaModel) Create(sewa entities.Sewa) bool {
    result, err := s.conn.Exec("INSERT INTO sewa (id_customer, id_motor, tanggal_sewa, tanggal_kembali) VALUES (?,?,?,?)", 
        sewa.CustomerId, sewa.MotorId, sewa.TanggalSewa, sewa.TanggalKembali)

    if err != nil {
        fmt.Println(err)
        return false 
    }

    lastInsertId, _ := result.LastInsertId()
    return lastInsertId > 0
}


func (s *SewaModel) FindAll() ([]entities.Sewa, error) {
	rows, err := s.conn.Query("SELECT sewa.id_sewa, customer.nama_lengkap, motor.tipe, sewa.tanggal_sewa, sewa.tanggal_kembali FROM sewa JOIN customer ON sewa.id_customer = customer.id_customer JOIN motor ON sewa.id_motor = motor.id_motor")
	if err != nil {
		return []entities.Sewa{}, err
	}
	defer rows.Close()

	var dataSewa []entities.Sewa
	for rows.Next() {
		var sewa entities.Sewa
		err := rows.Scan(&sewa.Id, &sewa.CustomerName, &sewa.TipeMotor, &sewa.TanggalSewa, &sewa.TanggalKembali)
		if err != nil {
			return []entities.Sewa{}, err
		}
		dataSewa = append(dataSewa, sewa)
	}
	return dataSewa, nil
}

func (s *SewaModel) Delete(id int64) {
	s.conn.Exec("delete from sewa where id_sewa = ?", id)
}