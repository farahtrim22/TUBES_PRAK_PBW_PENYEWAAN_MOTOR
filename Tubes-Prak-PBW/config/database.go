package config

import (
	"database/sql"
)

func DBconnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := ""
	dbName := "tubes_prak_pbw"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	return db, err
}
