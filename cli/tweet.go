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
	err := twc.Write(twdata)

	if err != nil {
		log.Println("fail write tweet")
		log.Println(err)
	} else {
		fmt.Println("success write")
	}

}
func (twc tweetCli) DeleteTweet(userId string) {
	err := twc.DeleteAllTweetFromUser(userId)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("success delete")
}
