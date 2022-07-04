package db

import (
	"github.com/go-redis/redis/v9"
)

type RedisClient struct {
	redisCl *redis.Client
}

func GetRedisClient() {

}
