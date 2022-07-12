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

	router := GetRouter(sfapi)

	router.Run(":8080")
}

func getRedisStore() redis.Store {
	pwd := []byte("secret")
	size := 10
	redisStore, _ := redis.NewStore(size, "tcp", "localhost:6379", "", pwd)
	return redisStore
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
