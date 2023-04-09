package apis

import (
	"net/http"
	"time"

	"github.com/NestorNeo/ades/modelo"
	"github.com/gin-gonic/gin"
)

func RegistrarAsistencia(c *gin.Context) {

	rfid, ok := c.GetQuery("rfid")

	if !ok {
		c.String(http.StatusBadRequest, "formato usa /sensor/asistencia?rfid=")
		return
	}

	db, err := helperGetDB(c)

	if err != nil {
		c.String(http.StatusInternalServerError, "Lo siento no ando jalando")
		return
	}

	estudiante := modelo.Estudiante{}

	db.Where("rf_id = ?", rfid).First(&estudiante)

	if estudiante.RFID == "" {
		c.String(http.StatusBadRequest, "NO esta el fantasma")
		return
	}

	asistencia := modelo.Asistencia{
		Fecha:        time.Now(),
		Presente:     true,
		EstudianteID: estudiante.ID,
	}

	db.Create(&asistencia)

	c.String(http.StatusAccepted, "welcome %d %s", estudiante.ID, asistencia.Fecha)

}
