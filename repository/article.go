package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sf-duplicate/request"
	"time"

	"github.com/go-redis/redis/v9"
)

// title: "Connecting Redis server with django"
// link: https://stackoverflow.com
// poster: user:832
// time: 1331355699.33
// votes: 250
type article struct {
	poster string  `redis:"poster"`
	title  string  `redis:"title"`
	link   string  `redis:"link"`
	time   int64   `redis:"time"`
	votes  float64 `redis:"time"`
}

type ArticleRepo struct {
	rcl *redis.Client
}

var ctx = context.TODO()

func NewArticleRepo(redCl *redis.Client) *ArticleRepo {
	return &ArticleRepo{redCl}
}

func (ar *ArticleRepo) orderByTime(time int64, articleHKey string) error {
	// ZADD time:
	fmt.Printf("add time %s \n", articleHKey)
	err := ar.rcl.ZAdd(ctx, "time:", redis.Z{
		Score:  float64(time),
		Member: articleHKey,
	}).Err()

	if err != nil {
		return fmt.Errorf("ZAdd time: has been Failed: %w", err)
	}
	return nil
}
func (ar *ArticleRepo) orderByScore(votes float64, articleHKey string) error {
	// ZADD score:
	// fmt.Printf("add score %s \n", articleHKey)
	err := ar.rcl.ZAdd(ctx, "score:", redis.Z{
		Score:  votes,
		Member: articleHKey,
	}).Err()

	if err != nil {
		return fmt.Errorf("ZAdd score: has been Failed: %w", err)
	}

	return nil
}

// contek saja vote yang ada pada buku, sebaiknya jangan buat rancang-data yang lain dulu
func (ar *ArticleRepo) Post(arReq request.ArticleRequest) error {
	var err error
	// generate article_id
	articleId := ar.rcl.Incr(ctx, "article:")
	articleHKey := fmt.Sprintf("article:%d", articleId.Val())

	// SADD voted = {votekey:score}
	votedSKey := fmt.Sprintf("voted:%d", articleId.Val())

	isdup, err := ar.rcl.SAdd(ctx, votedSKey, arReq.Votes).Result()

	if err != nil {
		return fmt.Errorf("SAdd Voted has been Failed: %w", err)
	}

	if isdup == 0 {
		return errors.New("user has voted (duplicated)")
	}

	errs := make(chan error, 1)

	go func() {
		errs <- ar.orderByScore(arReq.Votes, articleHKey)
	}()

	go func() {
		errs <- ar.orderByTime(arReq.Time, articleHKey)
	}()

	// HSET article
	isdup, err = ar.rcl.HSet(ctx, articleHKey, map[string]interface{}{
		"poster": arReq.Poster,
		"title":  arReq.Title,
		"link":   arReq.Link,
		"time":   time.Now().Unix(),
		"votes":  arReq.Votes,
	}).Result()

	if err != nil {
		return fmt.Errorf("HSet Article has been Failed: %w", err)
	}

	if isdup == 0 {
		return errors.New("article has been duplicated")
	}

	fmt.Printf("Article Id: %s : ", articleHKey)

	return <-errs
}

func (ar *ArticleRepo) Vote(zkey, member string, score int) {
	// score, members
	err := ar.rcl.ZAdd(ctx, zkey, redis.Z{
		Score:  float64(score),
		Member: member,
	}).Err()

	if err != nil {
		log.Println(err)
	}
	fmt.Println("success voted")
}

func (ar *ArticleRepo) GetArticles(order string) ([]article, error) {
	// SMEMBERS
	res, err := ar.rcl.SMembers(ctx, order).Result()
	if err != nil {
		return nil, err
	}
	var articles []article
	var art article
	// harus melibatkan 2 set, apakah pakai INTERSET?
	for _, arid := range res {
		res := ar.rcl.HGetAll(ctx, arid)
		err := res.Scan(&art)
		if err != nil {
			return nil, err
		}
		articles = append(articles, art)
	}

	return articles, nil
}

func (ar *ArticleRepo) Read() {
}
