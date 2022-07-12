package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetRouter(sfapi *SfApi) *gin.Engine {
	router := gin.Default()
	router.Use(sessions.Sessions("sfsession", getRedisStore()))

	router.GET("/", rootHandler)
	router.POST("/login", Login)
	router.POST("/logout", Logout)

	user := router.Group("/v1/user")
	user.Use(SessionAuthentication())

	{
		user.GET("/home", sfapi.listArticles)
	}
	return router
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "welcome API",
	})
}
