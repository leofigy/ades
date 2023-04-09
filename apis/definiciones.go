package apis

import "time"

type UIAsistenciaResult struct {
	ID              int
	Nombre          string
	ApellidoMaterno string
	ApellidoPaterno string
	RFID            string
	Fecha           time.Time
}

func (p UIAsistenciaResult) Alumno() string {
	return p.ApellidoPaterno + " " + p.ApellidoMaterno + p.Nombre
}
