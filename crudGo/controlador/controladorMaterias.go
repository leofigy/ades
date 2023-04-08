package controlador

import (

	//Envio y recepción de información
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
func CargarAgregarMaterias(w http.ResponseWriter, r *http.Request) {

	platillas.ExecuteTemplate(w, "crear", modelo.RecuperarMaterias())
}

//Insertar

func InsertarMateria2(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Un pasito mas para el Dominio Total del Mundo >:D ")
	if r.Method == "POST" {

		nombre := r.FormValue("nombre")
		aula := r.FormValue("aula")

		modelo.InsertarMateria(nombre, aula)
		http.Redirect(w, r, "/", 301)
	}
}

//******************************************************************
//******************************************************************
//***************Elminar una materia********************************
//******************************************************************
//******************************************************************

func CBorrarMateria(w http.ResponseWriter, r *http.Request) {

	idMateria := r.URL.Query().Get("id")

	modelo.MBorrarMateria(idMateria)

	http.Redirect(w, r, "/", 301)

}
