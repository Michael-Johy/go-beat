package milddle

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Time() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		elapsed := time.Since(t)

		log.Print(c.FullPath(), " ", elapsed.Milliseconds())

	}
}
