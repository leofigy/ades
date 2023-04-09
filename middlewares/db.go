package middlewares

import (
	"fmt"
	"log"

	"github.com/NestorNeo/ades/contracts"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GuidMiddleware(db contracts.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("okay testing again pal")
		uuid := uuid.New()
		c.Set("uuid", uuid.String())
		c.Set("db", db)
		fmt.Printf("The request with uuid %s is started \n", uuid)
		c.Next()
		fmt.Printf("The request with uuid %s is served \n", uuid)
	}
}
