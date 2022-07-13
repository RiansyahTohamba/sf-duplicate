package handler

import (
	"log"
	"sf-duplicate/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	arRepo  *repository.ArticleRepo
	usrRepo *repository.UserRepository
}

func NewArticleHandler(arRepo *repository.ArticleRepo, usrRepo *repository.UserRepository) *ArticleHandler {
	return &ArticleHandler{arRepo, usrRepo}
}

// recently view disimpan diuser atau di article?
// di user saja, lebih make-sense
func (arh *ArticleHandler) Detail(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get("user_id").(string)
	arh.usrRepo.WriteRecentlyView(userId)
}
func (arh *ArticleHandler) List(ctx *gin.Context) {
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
