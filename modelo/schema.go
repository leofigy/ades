package modelo

import (
	"time"

	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	Nombre          string
	ApellidoMaterno string
	ApellidoPaterno string
	Telefono        string
}

type Usuario struct {
	gorm.Model
	Password string
	Tipo     string
}

type Academico struct {
	gorm.Model
	Nombre          string
	ApellidoMaterno string
	ApellidoPaterno string
}

type Asistencia struct {
	gorm.Model
	Fecha        time.Time
	Presente     bool
	EstudianteID uint
}

type Rendimiento struct {
	gorm.Model
	Tarea        string
	Status       string
	Score        float64
	EstudianteID uint
}

type FaltaAdministrativa struct {
	gorm.Model
	Fecha        time.Time
	Razon        string
	EstudianteID uint
}
