package modelo

import (
	//Librerias para la conexion con la base de datos

	"database/sql"

	//IMPORTANTE ----previamente se descarga go get -u github.com/go-sql-driver/mysql---
	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	driver := "mysql"
	usuario := "root"
	pass := ""
	name := "felix"

	conexion, err := sql.Open(driver, usuario+":"+pass+"@tcp(127.0.0.1)/"+name)
	if err != nil {
		panic(err.Error())
	}
	return conexion

}
