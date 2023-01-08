package infrastructure

import (
	"net/http"
	"os"
	"strings"
	"time"

	account "github.com/kory-jp/vue_go/api/interfaces/controllers/account"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	accountController := account.NewAccountController(NewSqlHandler())
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("ALLOWED_ORIGINS"), " "),
		AllowMethods:     strings.Split(os.Getenv("ALLOWED_METHODS"), " "),
		AllowHeaders:     strings.Split(os.Getenv("ALLOWED_HEADERS"), " "),
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.POST("/register", func(c *gin.Context) {
		accountController.Create(c.Writer, c.Request)
	})
	r.Run(":8000")
}
