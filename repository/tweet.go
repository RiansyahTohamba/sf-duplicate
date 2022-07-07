package repository

import (
	"errors"
	"math/rand"
	"sf-duplicate/request"
	"time"

	"github.com/go-redis/redis/v9"
)

type Tweet struct {
	Id         int    `redis:"id"`
	Message    string `redis:"msg"`
	TimePosted int64  `redis:"time"`
	Retweet    int    `redis:"retweet"`
	Love       int    `redis:"love"`
}

type TweetRepo struct {
	rcl *redis.Client
}

func NewTweetRepo(rcl *redis.Client) *TweetRepo {
	return &TweetRepo{rcl}
}
func (twr *TweetRepo) SortByLove() {

}

// RT = retweet
func (twr *TweetRepo) SortByRT() []Tweet {
	return []Tweet{}
}

func (twr *TweetRepo) Write(twReq request.TweetRequest) error {
	// rq: how to generate id?
	id := rand.Intn(3000)

	isdup, err := twr.rcl.HSet(ctx, twReq.UserId, map[string]interface{}{
		"id":   id,
		"msg":  twReq.Message,
		"time": time.Now().Unix(),
	}).Result()

	if isdup == 0 {
		return errors.New("tweet has duplicated")
	}
	return err
}

func (twr *TweetRepo) Read(tweetId string) (Tweet, error) {
	// HGetAll indeed ambil semua sub key yg ada pada hash based on hashkey

	res := twr.rcl.HGetAll(ctx, tweetId)

	var twt Tweet
	err := res.Scan(&twt)

	if err != nil {
		return Tweet{}, err
	}

	return twt, nil
}
func (twr *TweetRepo) DeleteAllTweetFromUser(userId string) error {
	// Id         string `redis:"id"`
	// Message    string `redis:"msg"`
	// TimePosted int64  `redis:"time"`
	// 3 subkey nya harus dihapus satu persatu
	// hdel hash-key sub-key1
	// hdel hash-key sub-key2
	// hdel hash-key sub-key3
	// apakah ada cara yang praktis?

	columns := []string{"id", "msg", "time"}
	for _, v := range columns {
		err := twr.rcl.HDel(ctx, userId, v).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadMsg(tweetId string) {

}
