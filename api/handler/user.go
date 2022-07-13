package handler

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	// token & email didapat darimana?

	email := "mriansyah93@gmail.com"

	// userid akan dipakai untuk mencatat recently view
	// userId :=
	session.Set("userId", userId)
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

func generateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
