package cli

import (
	"fmt"
	"log"
	"sf-duplicate/repository"
	"sf-duplicate/request"
)

type tweetCli struct {
	repository.TweetRepo
}

func NewTweetCli(twr repository.TweetRepo) *tweetCli {
	return &tweetCli{twr}
}

func (twc tweetCli) ReadByUser(userId string) {
	fmt.Printf("all tweets from %s \n", userId)
	fmt.Println(twc.Read(userId))
}

func (twc tweetCli) WriteTweet(msg, userId string) {
	twdata := request.TweetRequest{UserId: userId, Message: msg}
	ok, err := twc.Write(twdata)
	// fmt.Println(ok)
	colsLen := 3

	if ok != int64(colsLen) {
		fmt.Println(ok)
		fmt.Println("gagal input")
	} else {
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("success write")

		}
	}

}
func (twc tweetCli) DeleteTweet(userId string) {
	err := twc.DeleteAllTweetFromUser(userId)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("success delete")
}
