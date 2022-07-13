package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RedisClient struct {
	*redis.Client
}

func GetRedisClient() *RedisClient {
	var once sync.Once
	var rcl *RedisClient
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		rcl = &RedisClient{client}
	})
	_, err := rcl.Ping(context.TODO()).Result()
	if err != nil {
		fmt.Println("running redis-server --daemonize yes")
		log.Fatalf("Could not connect to redis %v", err)
	}

	return rcl

}

func GetSqliteClient() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db/sfduplicate.db"), &gorm.Config{})
	return db, err
}
