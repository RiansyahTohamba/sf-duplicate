package main

import (
	"sf-duplicate/cli"
	"sf-duplicate/db"
	"sf-duplicate/repository"
)

func main() {
	client := db.GetRedisClient()
	twr := repository.NewTweetRepo(client.Client)
	twc := cli.NewTweetCli(*twr)

	// todo: marshall json.Body to struct
	user1 := "user:1"
	user2 := "user:29384"

	twc.WriteTweet("learn GO!", user1)
	twc.WriteTweet("semangat guys", user2)

	twc.ReadByUser(user1)

	// kalau sub-key nya lebih dari satu?
	twc.ReadByUser(user2)

	// userid dan tweet yang mana?
	// twc.DeleteTweet(user1)
	// twc.ReadByUser(user1)

}
