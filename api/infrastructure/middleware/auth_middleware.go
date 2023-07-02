package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kory-jp/vue_go/api/infrastructure/auth"
)

func AuthMiddleware(j *auth.JWTer) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		log.Print("AuthMiddleware")

		if err := j.FillContxet(ctx); err != nil {
			log.Print(err.Error())
			// TODO: エラーの返し方を考える
			return
		}
		ctx.Next()
	})
}
