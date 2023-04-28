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

	production := len(DBCredentials) > 0
	provider := modelo.NewProvider("datos.db", production, DBCredentials)

	_, err := provider.GetDB()

	if err != nil {
		log.Panic(err)
	}

	router := apis.GetAPIs(&provider)

	router.Run("0.0.0.0:8080")
}
