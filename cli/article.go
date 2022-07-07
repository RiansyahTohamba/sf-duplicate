package cli

import "sf-duplicate/repository"

type articleCli struct {
	repository.ArticleRepo
}

func NewArticleCli(arRepo repository.ArticleRepo) *articleCli {
	return &articleCli{arRepo}
}
