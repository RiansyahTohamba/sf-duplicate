package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (api *SfApi) listArticles(ctx *gin.Context) {
	articles, err := api.arRepo.GetArticles("time:", 10)

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
