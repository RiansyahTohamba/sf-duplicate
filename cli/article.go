package cli

import (
	"fmt"
	"log"
	"sf-duplicate/repository"
)

type articleCli struct {
	repository.ArticleRepo
}

func NewArticleCli(arRepo repository.ArticleRepo) *articleCli {
	return &articleCli{arRepo}
}

// siapa saja yang vote?
// input=
func (arc articleCli) getUserVoted() {
	// harus ada user struct disini
	// err := arc.Write(twdata)

}

// show me most scored article
func (arc articleCli) PrintMostScoredArticle() {
	res, err := arc.GetArticles("score:")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("most scored article")
	for _, v := range res {
		fmt.Println(v)
	}
}

// show me most recent article
func (arc articleCli) getRecentArticle() {

}
