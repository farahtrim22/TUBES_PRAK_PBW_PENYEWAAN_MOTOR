package models

import (
	"Tubes_PBW/config"
	"Tubes_PBW/entities"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MotorModel struct {
	conn *sql.DB
}

func NewMotorModel() *MotorModel {
	conn, err := config.DBconnection()
	if err != nil {
		panic(err)
	}

	return &MotorModel{
		conn: conn,
	}
}

func (m *MotorModel) FindAll() ([]entities.Motor, error) {
	rows, err := m.conn.Query("select * from motor")
	if err != nil {
		return []entities.Motor{}, err
	}
	defer rows.Close()

	var dataMotor []entities.Motor
	for rows.Next() {
		var motor entities.Motor
		rows.Scan(&motor.Id, &motor.Merek, &motor.Tipe, &motor.JenisMotor, &motor.TahunProduksi, &motor.Warna, &motor.Stok, &motor.Harga)
		dataMotor = append(dataMotor, motor)
	}
	return dataMotor, nil
}

func (m *MotorModel) Create(motor entities.Motor) bool {
	result, err := m.conn.Exec("INSERT INTO motor (merek, tipe, jenis_motor, tahun_produksi, warna, stok, harga) VALUES (?,?,?,?,?,?,?)",
		motor.Merek, motor.Tipe, motor.JenisMotor, motor.TahunProduksi, motor.Warna, motor.Stok, motor.Harga)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0

}

func (m *MotorModel) Find(id int64, motor *entities.Motor) error {
	return m.conn.QueryRow("select * from motor where id_motor = ?", id).Scan(&motor.Id, &motor.Merek, &motor.Tipe, &motor.JenisMotor, &motor.TahunProduksi, &motor.Warna, &motor.Stok, &motor.Harga)

}

func (m *MotorModel) Update(motor entities.Motor) error {
	_, err := m.conn.Exec(
		"update motor set merek = ?, tipe = ?, jenis_motor = ?, tahun_produksi = ?, warna = ?, stok = ?, harga = ? where id_motor = ?",
		motor.Merek, motor.Tipe, motor.JenisMotor, motor.TahunProduksi, motor.Warna, motor.Stok, motor.Harga, motor.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *MotorModel) Delete(id int64) {
	m.conn.Exec("delete from motor where id_motor = ?", id)
}

func (m *MotorModel) UpdateStok(motorID int64) error {
	_, err := m.conn.Exec("UPDATE motor SET stok = stok - 1 WHERE id_motor = ?", motorID)
	if err != nil {
		return err
	}
	return nil
}

