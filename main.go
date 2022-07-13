package main

import (
	"fmt"
	"log"
	"sf-duplicate/api"
	"sf-duplicate/cli"
	"sf-duplicate/db"
	"sf-duplicate/repository"
)

// JIKA CLI recently view sudah beres, lanjut ke Rest api based.

func main() {
	client := db.GetRedisClient()
	// sfExampleCli(client)
	// exampleApi(client)
	readArticlesCli(client)
}

func exampleApi(client *db.RedisClient) {
	fmt.Println("API example")

	orm, err := db.GetSqliteClient()
	if err != nil {
		log.Println(err)
	}

	usrRepo := repository.NewUserRepo(client.Client, orm)

	arRepo := repository.NewArticleRepo(client.Client)
	api.StartRouter(arRepo, usrRepo)
}

func readArticlesCli(client *db.RedisClient) {
	orm, err := db.GetSqliteClient()
	if err != nil {
		log.Println(err)
	}
	usrRepo := repository.NewUserRepo(client.Client, orm)
	// seedRecenltyView(usrRepo)
	userIdSession := "user:3"
	ids, err := usrRepo.GetRecentlyViews(userIdSession)
	if err != nil {
		log.Println(err)
	}
	for _, v := range ids {
		fmt.Println(v)
	}
}

func seedRecenltyView(usrRepo *repository.UserRepository) {
	// userid=3 read 3 articles
	userIdSession := "user:3"
	usrRepo.AddRecentlyView(userIdSession, 22)
	usrRepo.AddRecentlyView(userIdSession, 12)
	err := usrRepo.AddRecentlyView(userIdSession, 32)

	if err != nil {
		log.Println(err)
	}
}

func sfExampleCli(client *db.RedisClient) {
	arRepo := repository.NewArticleRepo(client.Client)
	arCli := cli.NewArticleCli(*arRepo)

	// post article or seed data
	seedData(arCli)
	// siapa saja yang vote?

	// show me most scored article
	// arCli.PrintMostScoredArticle()
	// show me most recent article
	// arCli.PrintRecentArticle()

}

func seedData(arCli *cli.ArticleCli) {
	postArticle1(arCli)
	postArticle2(arCli)
	postArticle3(arCli)
}
func postArticle1(arCli *cli.ArticleCli) {
	title := "How to specify go-redis expires"
	link := "https://stackoverflow.com"
	poster := "user:832"

	time := 1331344699
	votes := 528

	arCli.PostArticle(title, link, poster, int64(time), float64(votes))
}

func postArticle2(arCli *cli.ArticleCli) {
	title := "Connecting Redis server with django"
	link := "https://stackoverflow.com"
	poster := "user:832"
	time := 1331355610
	votes := 250

	arCli.PostArticle(title, link, poster, int64(time), float64(votes))
}

func postArticle3(arCli *cli.ArticleCli) {
	title := "How to migrate from RabbitMQ to REDIS in.net"
	link := "https://stackoverflow.com"
	poster := "user:832"
	time := 1331382699
	votes := 234

	arCli.PostArticle(title, link, poster, int64(time), float64(votes))
}

func tweetExample(client *db.RedisClient) {
	twr := repository.NewTweetRepo(client.Client)
	twc := cli.NewTweetCli(*twr)

	// todo: marshall json.Body to struct
	user1 := "user:1"
	user2 := "user:29384"
	// pesan twitter tidak bisa duplicate
	twc.WriteTweet("learn GO!", user1)
	twc.WriteTweet("semangat guys", user2)

	twc.ReadByUser(user1)

	// kalau sub-key nya lebih dari satu?
	twc.ReadByUser(user2)

	// userid dan tweet yang mana?
	// twc.DeleteTweet(user1)
	// twc.ReadByUser(user1)
}
