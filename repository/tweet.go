package repository

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v9"
)

type tweet struct {
	Message    string
	TimePosted time.Time
}

type TweetRepo struct {
	rcl *redis.Client
}

func (twr *TweetRepo) Write(tw tweet, key string) error {
	// rq: how to generate id?
	id := rand.Intn(1000)

	tweetId := fmt.Sprintf("tweet:%d", id)

	err := twr.rcl.HSet(ctx, tweetId, map[string]interface{}{
		"msg":  tw.Message,
		"time": tw.TimePosted.Unix(),
	}).Err()

	return err
}

func (twr *TweetRepo) Read(tweetId string) (tweet, error) {
	// HGetAll indeed ambil semua sub key yg ada pada hash based on hashkey
	// tw.rcl.HGetAll(ctx, tweetId)

	res := twr.rcl.HGetAll(ctx, tweetId)
	// Expect(res.Err()).NotTo(HaveOccurred())

	// type data struct {
	// 	Key1 string `redis:"key1"`
	// 	Key2 int    `redis:"key2"`
	// }
	var twt tweet
	err := res.Scan(&twt)

	if err != nil {
		return tweet{}, err
	}

	return twt, nil
}

func getMsg() {

}
