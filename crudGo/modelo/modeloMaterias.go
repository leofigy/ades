package modelo

import (

	//Libreria necesaria para ejecutar sentacias sql
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Estructura de materia_ Debe concotrdad con la tabla de la base de datos
type materia struct {
	Id     int
	Nombre string
	Aula   string
}

//Funcion para agregar una materia

func InsertarMateria(nombre, aula string) {

	conexionestab := conexionBD()

	insertarRegistro, err := conexionestab.Prepare("INSERT INTO materia (nombre, aula) VALUES (?,?)")

	if err != nil {
		panic(err.Error())

	}
	insertarRegistro.Exec(nombre, aula)

}

// Funcion para recuperar muchas materias
func RecuperarMaterias() []materia {

	conexionestab := conexionBD()

	registros, err := conexionestab.Query("SELECT * FROM materia")

	//Encontrar posibles errores
	if err != nil {
		panic(err.Error())
	}
	materias := materia{}
	arregloMaterias := []materia{}

	//Austar los datos recuperados a una array
	for registros.Next() {

		//Se debe especificar todos los datos de la tabla a trabajar
		var id int
		var nombre, aula string
		err = registros.Scan(&id, &nombre, &aula)
		if err != nil {
			panic(err.Error())
		}

		//Te permite trabajar solo con los datos que requieres
		materias.Id = id
		materias.Nombre = nombre
		materias.Aula = aula

		arregloMaterias = append(arregloMaterias, materias)

	}

	//retornar valores recuperados en un array de tipo estructura ( []materia)
	fmt.Print(arregloMaterias)
	return arregloMaterias

}

//Elimanr

func MBorrarMateria(idMateria string) {
	conexionestab := conexionBD()
	//insertarRegistro,err:= conexionestab.Prepare("INSERT INTO esudiante (idEsudiante, nombre, apellidoPaterno, apellidoMaterno, fechaNacimiento, grado, grupo, correo, rfid, historialMedico) VALUES ('3', 'Nestor', 'Caza', 'Rubias', '2015-03-23', '3', 'C', 'nestitor@gmail.com', 'zhvjwjvjw222', 'alergia a los gatos')")

	borrarRegistro, err := conexionestab.Prepare("DELETE FROM materia WHERE idmateria=?")

	if err != nil {
		panic(err.Error())

	}
	borrarRegistro.Exec(idMateria)
}
