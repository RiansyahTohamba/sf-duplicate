package request

type ArticleRequest struct {
	// sebagai hash key
	Poster string
	Title  string
	Link   string
	Time   int64
	Votes  float64
}

// title: "How to specify go-redis expires"
// link: https://stackoverflow.com
// poster: user:832
// time: 1331344699.33
// votes: 528
