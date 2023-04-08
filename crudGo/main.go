package main

import (
	"net/http"
	//Mostrar emnsajes en pantalla
	"log"

	"crudGo/controlador"
)

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", CrearEstudiante)

	//acciones de botones
	http.HandleFunc("/insertar", controlador.InsertarEstudiante)
	http.HandleFunc("/borrar", Borrar)

	log.Println("Servidor Corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Borrar(w http.ResponseWriter, r *http.Request) {

	controlador.CBorrarMateria(w, r)

}

func Inicio(w http.ResponseWriter, r *http.Request) {
	controlador.CargarInicio(w, r)

}

func CrearEstudiante(w http.ResponseWriter, r *http.Request) {
	controlador.CargarAgeragarEstudiantes(w, r)
}
