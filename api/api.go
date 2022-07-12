package api

import (
	"sf-duplicate/repository"

	"github.com/gin-contrib/sessions/redis"
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
