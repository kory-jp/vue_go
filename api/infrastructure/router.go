package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":8000")
}
