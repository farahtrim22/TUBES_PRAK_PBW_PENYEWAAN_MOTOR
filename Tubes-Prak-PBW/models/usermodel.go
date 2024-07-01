package models

import (
	"Tubes_PBW/config"
	"Tubes_PBW/entities"
	"database/sql"
)

type UserModel struct {
	conn *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBconnection()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		conn: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {
	row, err := u.conn.Query("select * from user where " + fieldName + " = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Username, &user.Password)
	}

	return nil
}

func (u UserModel) Create(user entities.User) (int64, error) {

	result, err := u.conn.Exec("insert into user (nama_lengkap, email, username, password) values(?,?,?,?)",
		user.NamaLengkap, user.Email, user.Username, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil

}