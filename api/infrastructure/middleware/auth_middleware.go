package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kory-jp/vue_go/api/infrastructure/auth"
)

func AuthMiddleware(j *auth.JWTer) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

		if err := j.FillContxet(ctx); err != nil {
			log.Print(err.Error())
			// TODO: エラーの返し方を考える
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	})
}
