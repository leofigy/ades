package controlador

import (

	//Envio y recepción de información
	"net/http"
	//Mostrar emnsajes en pantalla

	"text/template"
	//previamente se descarga go get -u github.com/go-sql-driver/mysql

	//previamente se descarga go get -u github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"

	"crudGo/modelo"
)

// Se declara la variable para cargar la plantilla
var platillas = template.Must(template.ParseGlob("platillas/*"))

// Cargar platilla con la informacion de la base de datos
func CargarInicio(w http.ResponseWriter, r *http.Request) {

	//Se ejecuta la platilla inicio pasandole los datos recuperados con la funcion ---  modelo.RecuperarMaterias()---
	platillas.ExecuteTemplate(w, "inicio", modelo.RecuperarEstudiantes())
}
