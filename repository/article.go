package repository

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/go-redis/redis/v9"
)

// title: "Connecting Redis server with django"
// link: https://stackoverflow.com
// poster: user:832
// time: 1331355699.33
// votes: 250
type article struct {
	title  string
	link   string
	poster string
	time   int64
}

type ArticleRepo struct {
	redCl *redis.Client
}

var ctx = context.TODO()

func NewArticleRepo(redCl *redis.Client) *ArticleRepo {
	return &ArticleRepo{redCl}
}

// hash-key: article:22
// title: "Connecting Redis server with django"
// link: https://stackoverflow.com
// poster: user:832
// time: 1331355699.33
// votes: 250

func (ar *ArticleRepo) Create() {
	// ar.redCl.Set()
	id := rand.Intn(1000)
	articleId := fmt.Sprintf("article:%d", id)
	article := map[string]article{}
	err := ar.redCl.HSet(ctx, articleId, article)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("article success created")
}

// hSet := client.HSet(ctx, "hash", "key", "hello")
// Expect(hSet.Err()).NotTo(HaveOccurred())

// hDel := client.HDel(ctx, "hash", "key")
// Expect(hDel.Err()).NotTo(HaveOccurred())
// Expect(hDel.Val()).To(Equal(int64(1)))

func (ar *ArticleRepo) Vote(zkey, member string, score int) {
	// score, members
	err := ar.redCl.ZAdd(ctx, zkey, redis.Z{
		Score:  float64(score),
		Member: member,
	}).Err()

	if err != nil {
		log.Println(err)
	}
	fmt.Println("success voted")
}

func (ar *ArticleRepo) Read() {

}

func (ar *ArticleRepo) FindAll() {

}
