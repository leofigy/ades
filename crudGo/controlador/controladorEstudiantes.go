package controlador

import (

	//Envio y recepción de información
	"fmt"
	"net/http"

	//Mostrar emnsajes en pantalla
	//previamente se descarga go get -u github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"

	"crudGo/modelo"
)

//------------------------------------------------------------------
//------------------------------------------------------------------
//--------------------Agregar una nueva materia---------------------
//------------------------------------------------------------------
//------------------------------------------------------------------

// Cargar platilla con la informacion de la base de datos
func CargarAgeragarEstudiantes(w http.ResponseWriter, r *http.Request) {

	platillas.ExecuteTemplate(w, "crear", modelo.RecuperarEstudiantes())
}

//Insertar

func InsertarEstudiante(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Un pasito mas para el Dominio Total del Mundo >:D ")
	fmt.Print("llegue a 1")
	if r.Method == "POST" {

		nombre := r.FormValue("nombre")
		apellidoPaterno := r.FormValue("apellidoPaterno")
		apellidoMaterno := r.FormValue("apellidoMaterno")

		fmt.Println("----_")
		fmt.Println(nombre)

		modelo.InsertarEstudiante(nombre, apellidoPaterno, apellidoMaterno)
		http.Redirect(w, r, "/", 301)
	}
}

//******************************************************************
//******************************************************************
//***************Elminar una materia********************************
//******************************************************************
//******************************************************************

func CBorrarEstudiante(w http.ResponseWriter, r *http.Request) {

	idEstudiante := r.URL.Query().Get("id")

	modelo.MBorrarEstudiante(idEstudiante)

	http.Redirect(w, r, "/", 301)

}
