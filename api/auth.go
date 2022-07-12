package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	// id & email didapat darimana?
	id := 122323
	email := "mriansyah93@gmail.com"

	session.Set("id", id)
	session.Set("email", email)
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully",
	})
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success logout",
	})
}
