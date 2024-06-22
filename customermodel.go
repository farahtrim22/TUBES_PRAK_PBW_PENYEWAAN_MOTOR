package models

import (
	"Tubes_PBW/config"
	"Tubes_PBW/entities"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerModel struct {
	conn *sql.DB
}

func NewCustomerModel() *CustomerModel {
	conn, err := config.DBconnection()
	if err != nil {
		panic(err)
	}

	return &CustomerModel{
		conn: conn, 
	}
}

func (c *CustomerModel) FindAll() ([]entities.Customer, error) {
	rows, err := c.conn.Query("select * from customer")
	if err != nil {
		return []entities.Customer{}, err
	} 
	defer rows.Close()

	var dataCustomer []entities.Customer
	for rows.Next() {
		var customer entities.Customer
		rows.Scan(&customer.Id, &customer.NamaLengkap, &customer.NIK, &customer.JenisKelamin, &customer.Alamat, &customer.NomorTelepon)

		if customer.JenisKelamin == "1" {
			customer.JenisKelamin = "Laki-Laki"
		} else {
			customer.JenisKelamin = "Perempuan"
		}

		dataCustomer = append(dataCustomer, customer)
	}
	return dataCustomer, nil
}

func (c *CustomerModel) Create(customer entities.Customer) bool {
	result, err := c.conn.Exec("INSERT INTO customer (nama_lengkap, nik, jenis_kelamin, alamat, nomor_telepon) VALUES (?,?,?,?,?)", 
	customer.NamaLengkap, customer.NIK, customer.JenisKelamin, customer.Alamat, customer.NomorTelepon)

	if err != nil {
		fmt.Println(err)
		return false 
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
	
}

func (c *CustomerModel) Find(id int64, customer *entities.Customer) error {
	return c.conn.QueryRow("select * from customer where id = ?", id).Scan(&customer.Id, &customer.NamaLengkap, &customer.NIK, &customer.JenisKelamin, &customer.Alamat, &customer.NomorTelepon)

}

func (c *CustomerModel) Update(customer entities.Customer) error {
	_, err := c.conn.Exec(
		"update customer set nama_lengkap = ?, nik = ?, jenis_kelamin = ?, alamat = ?, nomor_telepon = ? where id = ?",
		customer.NamaLengkap, customer.NIK, customer.JenisKelamin, customer.Alamat, customer.NomorTelepon, customer.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *CustomerModel) Delete(id int64) {
	c.conn.Exec("delete from customer where id = ?", id)
}
