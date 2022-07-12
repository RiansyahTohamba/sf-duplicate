package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func sessionAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		session := sessions.Default(ctx)
		sessionID := session.Get("userId")

		if sessionID == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "unauthorized",
			})
			ctx.Abort()
		}
		// check token disini?

	}
}
