package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetRouter(sfapi *SfApi) *gin.Engine {
	router := gin.Default()
	router.Use(sessions.Sessions("sfsession", getRedisStore()))

	router.GET("/", rootHandler)

	router.Use(sessions.Sessions("counter", getRedisStore()))

	router.GET("/incr", incrementHandler)

	router.POST("/login", Login)
	router.POST("/logout", Logout)

	user := router.Group("/api/v1/user")
	user.Use(sessionAuth())

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

// try redis session
func incrementHandler(ctx *gin.Context) {
	var counter int
	session := sessions.Default(ctx)
	val := session.Get("counter")

	if val == nil {
		counter = 0
	} else {
		counter = val.(int)
		counter++
	}
	session.Set("counter", counter)
	session.Save()

	ctx.JSON(200, gin.H{
		"counter": counter,
	})
}
