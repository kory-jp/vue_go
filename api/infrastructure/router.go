package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kory-jp/vue_go/api/infrastructure/auth"
	"github.com/kory-jp/vue_go/api/infrastructure/middleware"
	"github.com/kory-jp/vue_go/api/infrastructure/store"

	"github.com/gin-gonic/gin"
	"github.com/kory-jp/vue_go/api/interfaces/controllers"
	account_controller "github.com/kory-jp/vue_go/api/interfaces/controllers/account"
	auth_controller "github.com/kory-jp/vue_go/api/interfaces/controllers/auth"

	"github.com/gin-contrib/cors"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func (res *Response) setResp(c *gin.Context, handler func(c controllers.Context) (status int, message string, body interface{}, err error)) {
	status, mess, body, err := handler(c)
	if err != nil {
		if e, ok := err.(fmt.Formatter); ok {
			log.Printf("[ERROR]: %+v\n\n", e)
		}
	}
	response := &Response{status, mess, body}
	jsonData, _ := json.Marshal(response)
	resStr := string(jsonData)
	fmt.Fprintln(c.Writer, resStr)
}

func (res *Response) setLoginResp(c *gin.Context, jwter *auth.JWTer, handler func(c controllers.Context, jwter controllers.JWTer) (status int, message string, body interface{}, err error)) {
	status, mess, body, err := handler(c, jwter)
	if err != nil {
		if e, ok := err.(fmt.Formatter); ok {
			log.Printf("[ERROR]: %+v\n\n", e)
		}
	}
	response := &Response{status, mess, body}
	jsonData, _ := json.Marshal(response)
	resStr := string(jsonData)
	fmt.Fprintln(c.Writer, resStr)
}

func Init() {
	log.SetFlags(log.Ltime | log.Llongfile)
	kvs, err := store.NewKVS()
	if err != nil {
		log.Printf("[ERROR]: %+v", err)
	}
	jwter, err := auth.NewJWTer(kvs)
	if err != nil {
		log.Printf("[ERROR]: %+v", err)
	}
	log.Printf("jwter: %v", jwter)
	accountController := account_controller.NewAccountController(NewSqlHandler())
	loginController := auth_controller.NewAuthController(NewSqlHandler())

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
		new(Response).setResp(c, accountController.Create)
	})

	r.POST("/api/login", func(c *gin.Context) {
		new(Response).setLoginResp(c, jwter, loginController.Login)
	})

	authRouter := r.Group("api/v1").Use(middleware.AuthMiddleware(jwter))

	authRouter.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get Tasks!",
		})
	})

	r.Run(":8000")

}
