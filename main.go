package main

import (
	//Mostrar emnsajes en pantalla
	"flag"
	"log"

	"github.com/NestorNeo/ades/apis"
	"github.com/NestorNeo/ades/modelo"
)

var (
	DBCredentials string
	production    bool // false uses sqlite
)

func init() {

	flag.StringVar(
		&DBCredentials,
		"db",
		"",
		"specify the connection string for mariadb, if not this uses sqlite internally",
	)

}

func main() {

	flag.Parse()
	if len(DBCredentials) == 0 {
		production = false
	}

	provider := modelo.NewProvider("datos.db")

	_, err := provider.GetDB()

	if err != nil {
		log.Panic(err)
	}

	router := apis.GetAPIs(&provider)

	router.Run("0.0.0.0:8080")
}
