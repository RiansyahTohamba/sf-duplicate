package handler

import (
	"log"
	"sf-duplicate/repository"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	arRepo *repository.ArticleRepo
}

func NewArticleHandler(arRepo *repository.ArticleRepo) *ArticleHandler {
	return &ArticleHandler{arRepo}
}

func (arh *ArticleHandler) ListArticles(ctx *gin.Context) {
	articles, err := arh.arRepo.GetArticles("time:", 10)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "error happened, call the programmer!",
		})
		log.Println(err)
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
		"data":    articles,
	})
}
