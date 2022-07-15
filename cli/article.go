package cli

import (
	"fmt"
	"log"
	"sf-duplicate/repository"
	"sf-duplicate/request"
)

var articlePerPage int64 = 5

type ArticleCli struct {
	repository.ArticleRepo
}

func NewArticleCli(arRepo repository.ArticleRepo) *ArticleCli {
	return &ArticleCli{arRepo}
}

// func (twc tweetCli) WriteTweet(msg, userId string) {
// 	twdata := request.TweetRequest{UserId: userId, Message: msg}
// 	err := twc.Write(twdata)

func (arc ArticleCli) PostArticle(title, link, poster string, time int64, votes float64) {
	arData := request.ArticleRequest{Title: title, Link: link, Poster: poster, Time: time, Votes: votes}
	err := arc.Post(arData)

	if err != nil {
		log.Println("fail post article")
		log.Println(err)
	} else {
		fmt.Println("success post article")
	}
}

// siapa saja yang vote?
// input=
func (arc ArticleCli) PrintUserVoted() {
	// harus ada user struct disini
	// err := arc.Write(twdata)

}

// show me most scored article
func (arc ArticleCli) PrintMostScoredArticle(page int64) {
	res, err := arc.GetArticles("score:", page, articlePerPage)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("most scored article")
	for _, v := range res {
		fmt.Println(v)
	}
}

// show me most recent article
func (arc ArticleCli) PrintRecentArticle(page int64) {

	res, err := arc.GetArticles("time:", page, articlePerPage)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("page %d \n", page)
	for _, v := range res {
		fmt.Println(v.Title)
	}

	fmt.Printf("============= \n")

}
