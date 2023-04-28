package modelo

// for testing in sqlite (all included)

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Estudiante struct {
	gorm.Model
	// foreign keys
	MiTutorRefer    int
	Nombre          string
	ApellidoPaterno string
	ApellidoMaterno string
	RFID            string
	MiTutor         Tutor     `gorm:"foreignKey:MiTutorRefer"`
	Materias        []Materia `gorm:"many2many:estudiante_materias;"`
	Asistencias     []Asistencia
	Reportes        []FaltaAdministrativa
	Rendimiento     []Rendimiento
}

type Materia struct {
	gorm.Model
	Nombre      string
	Aula        string
	Hora        string
	Dias        string
	Estudiantes []*Estudiante `gorm:"many2many:estudiante_materias;"`
}

type DBProvider struct {
	db          *gorm.DB
	location    string
	production  bool
	credentials string
}

func NewProvider(location string, production bool, credentials string) DBProvider {
	return DBProvider{
		nil, location, production, credentials,
	}
}

func (provider *DBProvider) GetDB() (*gorm.DB, error) {
	if provider.db != nil {
		return provider.db, nil
	}
	err := provider.InitModels()
	return provider.db, err
}

// database
// db location
func (provider *DBProvider) InitModels() error {

	var (
		db  *gorm.DB
		err error
	)

	if provider.production {
		log.Println("-------- production -------")
		db, err = gorm.Open(
			mysql.Open(provider.credentials), &gorm.Config{},
		)
	} else {
		log.Println("-------- starting in testing mode -------")
		db, err = gorm.Open(sqlite.Open(provider.location), &gorm.Config{})
	}

	if err != nil {
		return nil
	}

	db.AutoMigrate(
		&Estudiante{},
		&Asistencia{},
		&Rendimiento{},
		&Tutor{},
		&FaltaAdministrativa{},
		&Estudiante{},
		&Materia{},
	)

	if provider.db == nil {
		provider.db = db
	}

	return nil
}
