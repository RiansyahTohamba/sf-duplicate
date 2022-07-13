package api

import (
	"sf-duplicate/api/handler"
	"sf-duplicate/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func StartRouter(arRepo *repository.ArticleRepo) {
	arHandler := handler.NewArticleHandler(arRepo)
	usrHandler := handler.NewUserHandler(arRepo)

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
		user.GET("/home", arHandler.listArticles)
	}
	router.Run(":8080")
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "welcome API",
	})
}

func getRedisStore() redis.Store {
	pwd := []byte("secret")
	size := 10
	redisStore, _ := redis.NewStore(size, "tcp", "localhost:6379", "", pwd)
	return redisStore
}

// try redis session
func incrementHandler(ctx *gin.Context) {
	var counter int
	session := sessions.Default(ctx)

	token := session.Get("token")
	val := session.Get("counter")

	if val == nil {
		counter = 0
	} else {
		counter = val.(int)
		counter++
	}
	session.Set("counter", counter)
	session.Set("token", generateSecureToken(10))
	session.Save()

	ctx.JSON(200, gin.H{
		"counter": counter,
		"token":   token,
	})
}
