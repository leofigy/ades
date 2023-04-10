package apis

import (
	"errors"
	"log"
	"net/http"

	"github.com/NestorNeo/ades/contracts"
	"github.com/NestorNeo/ades/middlewares"
	"github.com/NestorNeo/ades/modelo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAPIs(provider contracts.Provider) *gin.Engine {
	r := gin.Default()

	// setting pre-sets
	r.Use(
		middlewares.GuidMiddleware(provider),
	)

	r.LoadHTMLGlob(
		"plantillas/*",
	)

	r.NoRoute(UINotFound)

	// Only UI ---------------------------------
	r.GET("/", Inicio)
	r.POST("/", Inicio)

	r.GET("/estudiantes/agregar", UICrearEstudiante)
	r.GET("/tareas/agregar", UICrearTarea)
	r.GET("/asistencias", UIAsistencias)
	r.GET("/tutores", UITutores)
	r.POST("/busqueda", UIBusqueda)

	// backend logic with forms goes here
	r.POST("/estudiantes", CrearEstudiantes)
	r.POST("/tutores", CrearTutores)

	// Internal APIs (just (REST))
	r.GET("/sensor/asistencia", RegistrarAsistencia)

	return r
}

func Inicio(c *gin.Context) {
	// getting info from context
	db, err := helperGetDB(c)

	if err != nil {
		UIError(
			http.StatusInternalServerError,
			err,
			c,
		)
		return
	}

	// temporal datos hardcodeados
	estudiantes := []modelo.Estudiante{}
	// get all users
	result := db.Find(&estudiantes)

	if result.Error != nil {
		UIError(
			http.StatusInternalServerError,
			err,
			c,
		)
		return
	}

	c.HTML(
		http.StatusOK,
		"inicio",
		gin.H{
			"Estudiantes": estudiantes,
		},
	)
}

func UITutores(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"creartutor",
		gin.H{},
	)
}

func UICrearTarea(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"crear",
		gin.H{})
}

func UICrearEstudiante(c *gin.Context) {
	db, err := helperGetDB(c)

	if err != nil {
		UIError(
			http.StatusInternalServerError,
			err,
			c,
		)
		return
	}
	tutors := []modelo.Tutor{}
	result := db.Find(&tutors)

	if result.Error != nil {
		UIError(
			http.StatusInternalServerError,
			err,
			c,
		)
		return
	}

	c.HTML(
		http.StatusOK,
		"crear",
		gin.H{
			"Padres": tutors,
		})
}

func UIBusqueda(c *gin.Context) {

	keyword := c.PostForm("keyword")
	modeloBusqueda := c.PostForm("modelo_busqueda")

	log.Println("HELLO ----- Looking for ", keyword, modeloBusqueda)

	db, err := helperGetDB(c)
	if err != nil {
		UIError(
			http.StatusInternalServerError,
			err,
			c,
		)
		return
	}

	// materia/estudiante/tutor
	switch modeloBusqueda {
	case "Estudiante":
		estudiantes := []modelo.Estudiante{}
		// take a look to
		// https://gorm.io/docs/query.html
		result := db.Preload("MiTutor").Where(
			"nombre LIKE ?", keyword,
		).Find(&estudiantes)

		if result.Error != nil {
			UIError(
				http.StatusInternalServerError,
				err,
				c,
			)
			return
		}

		c.HTML(
			http.StatusOK,
			"busquedastudent",
			gin.H{"Estudiantes": estudiantes})
		return

	default:
		UINotFound(c)
	}

}

func MateriaBorrar(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"crear",
		gin.H{})
}

func UIAsistencias(c *gin.Context) {
	// pending to validate
	db, err := helperGetDB(c)

	if err != nil {
		UIError(
			http.StatusInternalServerError,
			err,
			c,
		)
		return
	}

	asistencias := []UIAsistenciaResult{}

	db.Order("fecha desc").Table("asistencia").Select("estudiantes.id, estudiantes.nombre, estudiantes.apellido_materno, estudiantes.apellido_paterno ,estudiantes.rf_id", "asistencia.fecha").Joins("left join estudiantes on estudiantes.id = asistencia.estudiante_id").Scan(&asistencias)
	// ponganle paginacion xD es decir el boton de next
	log.Println(asistencias)
	c.HTML(
		http.StatusOK,
		"asistencias",
		gin.H{
			"Asistencias": asistencias,
		})
}

func UINotFound(c *gin.Context) {
	UIError(http.StatusNotFound, errors.New("no encontrado"), c)
}

func UIError(code int, err error, c *gin.Context) {
	c.HTML(
		code,
		"error",
		gin.H{
			"Code":  code,
			"Error": err,
		},
	)
}

func CrearEstudiantes(c *gin.Context) {

	// first validate input
	student := modelo.Estudiante{
		Nombre:          c.PostForm("nombre"),
		ApellidoPaterno: c.PostForm("ApellidoPaterno"),
		ApellidoMaterno: c.PostForm("ApellidoMaterno"),
		RFID:            c.PostForm("RFID"),
	}

	tutorId := c.PostForm("tutor")
	log.Println("este es el tutor", tutorId)

	// pending to validate
	db, err := helperGetDB(c)
	tutor := modelo.Tutor{}
	db.Find(&tutor, tutorId)

	log.Println("ENCONTRAMOS AL TUTOR ", tutor)

	student.MiTutor = tutor

	if err != nil {
		UIError(
			http.StatusInternalServerError,
			err,
			c,
		)
		return
	}

	log.Println("entry for db", student)
	db.Create(&student)
	log.Println("creando en la db pal un estudiante")
	c.Redirect(http.StatusPermanentRedirect, "/")

}

func CrearTutores(c *gin.Context) {
	tutor := modelo.Tutor{
		Nombre:          c.PostForm("nombre"),
		ApellidoPaterno: c.PostForm("ApellidoPaterno"),
		ApellidoMaterno: c.PostForm("ApellidoMaterno"),
	}

	db, err := helperGetDB(c)
	if err != nil {
		UIError(
			http.StatusInternalServerError,
			err,
			c,
		)
		return
	}

	log.Println("entry for tutor db", tutor)
	db.Create(&tutor)
	log.Println("creando en la db pal un tutor")
	c.Redirect(http.StatusPermanentRedirect, "/")
}

func helperGetDB(c *gin.Context) (*gorm.DB, error) {
	obj, ok := c.Get("db")

	if !ok {
		return nil, errors.New("no database available in context")
	}
	provider := obj.(contracts.Provider)
	return provider.GetDB()
}
