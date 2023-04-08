package modelo

import (

	//Libreria necesaria para ejecutar sentacias sql
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Estructura de materia_ Debe concotrdad con la tabla de la base de datos
type estudiante struct {
	IdEstudiante    int
	Nombre          string
	ApellidoPaterno string
	ApellidoMaterno string
}

//Funcion para agregar una materia

func InsertarEstudiante(nombre, apellidoPaterno, apellidoMaterno string) {
	fmt.Print("llegue a 2")
	conexionestab := conexionBD()

	fmt.Println("*****")
	fmt.Print(nombre)

	//llaveForeanaPapa := '1'

	//al agregar un elemnto siempre debe agregar la llave(S) foraneas en caso de tenerla en la tabla
	insertarRegistro, err := conexionestab.Prepare("INSERT INTO `esudiante` (`idEsudiante`, `nombre`, `apellidoPaterno`, `apellidoMaterno`, `fechaNacimiento`, `grado`, `grupo`, `correo`, `rfid`, `historialMedico`, `padreFamilia_idEsudiante`) VALUES (NULL, 'Marcos', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '1')")

	if err != nil {
		panic(err.Error())

	}
	insertarRegistro.Exec()

}

// Funcion para recuperar muchas materias
func RecuperarEstudiantes() []estudiante {

	conexionestab := conexionBD()

	registros, err := conexionestab.Query("SELECT * FROM esudiante")

	//Encontrar posibles errores
	if err != nil {
		panic(err.Error())
	}
	estudiantes := estudiante{}
	arregloEstudiantes := []estudiante{}

	//Austar los datos recuperados a una array
	for registros.Next() {
		var id int
		// Los Xn representan los datos que no estoy utilizando en mi estructura, devido a que no las utilizo
		var nombre, apellidoPaterno, apellidoMaterno, x1, x2, x3, x4, x5, x6, x7 string
		err = registros.Scan(&id, &nombre, &apellidoPaterno, &apellidoMaterno, &x1, &x2, &x3, &x4, &x5, &x6, &x7)
		if err != nil {
			panic(err.Error())
		}
		estudiantes.IdEstudiante = id
		estudiantes.Nombre = nombre
		estudiantes.ApellidoPaterno = apellidoPaterno
		estudiantes.ApellidoMaterno = apellidoMaterno

		arregloEstudiantes = append(arregloEstudiantes, estudiantes)

	}

	//retornar valores recuperados en un array de tipo estructura ( []materia)
	return arregloEstudiantes

}

//Elimanr

func MBorrarEstudiante(idMateria string) {
	conexionestab := conexionBD()
	//insertarRegistro,err:= conexionestab.Prepare("INSERT INTO esudiante (idEsudiante, nombre, apellidoPaterno, apellidoMaterno, fechaNacimiento, grado, grupo, correo, rfid, historialMedico) VALUES ('3', 'Nestor', 'Caza', 'Rubias', '2015-03-23', '3', 'C', 'nestitor@gmail.com', 'zhvjwjvjw222', 'alergia a los gatos')")

	borrarRegistro, err := conexionestab.Prepare("DELETE FROM materia WHERE idmateria=?")

	if err != nil {
		panic(err.Error())

	}
	borrarRegistro.Exec(idMateria)
}
