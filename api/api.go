package api

import (
	"net/http"
	"sf-duplicate/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type SfApi struct {
	arRepo *repository.ArticleRepo
}

func NewApi(arRepo *repository.ArticleRepo) *SfApi {
	return &SfApi{arRepo}
}

func StartRouter(arRepo *repository.ArticleRepo) {
	sfapi := NewApi(arRepo)

	router := gin.Default()

	secret := []byte("secret")
	size := 10

	store, _ := redis.NewStore(size, "tcp", "localhost:6379", "", secret)

	router.Use(sessions.Sessions("sfsession", store))

	router.POST("/login", Login)
	router.POST("/logout", Logout)
	user := router.Group("/v1/user")

	user.Use(SessionAuthentication())

	{
		user.GET("/home", sfapi.listArticles)
	}

	router.Run(":8080")
}

func SessionAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionID := session.Get("id")
		if sessionID == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "unauthorized",
			})
			ctx.Abort()
		}
	}
}
