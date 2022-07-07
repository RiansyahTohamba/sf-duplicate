package request

type ArticleRequest struct {
	Poster string
	Title  string
	Link   string
	Time   int64
	Votes  float64
}
