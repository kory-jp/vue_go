package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	account "github.com/kory-jp/vue_go/api/interfaces/controllers/account"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func (res *Response) setResp(w http.ResponseWriter, r *http.Request, handler func(r *http.Request) (status int, message string, body interface{}, err error)) {
	status, mess, body, err := handler(r)
	if err != nil {
		if e, ok := err.(fmt.Formatter); ok {
			log.Printf("[ERROR]: %+v\n\n", e)
		}
	}
	response := &Response{status, mess, body}
	jsonData, _ := json.Marshal(response)
	resStr := string(jsonData)
	fmt.Fprintln(w, resStr)
}

func Init() {
	log.SetFlags(log.Ltime | log.Llongfile)
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
	r.POST("/api/register", func(c *gin.Context) {
		new(Response).setResp(c.Writer, c.Request, accountController.Create)
	})
	r.Run(":8000")
}
