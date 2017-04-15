package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func getUser(client *twitter.Client) {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("enter username: @")
	username, _ := newcommand.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Println("user selected: " + username)

	fmt.Println("-----------------------")

	tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: username,
		Count:      200,
	})

	fmt.Println("word count")
	analyzeWords(tweets)

	analyzeDates(tweets)

	fmt.Println(" ")
	fmt.Print("first tweet analyzed: ")
	fmt.Println(tweets[len(tweets)-1].CreatedAt)
	fmt.Print("last tweet analyzed: ")
	fmt.Println(tweets[0].CreatedAt)

	fmt.Println(" ")
	fmt.Println("User @" + username + " analysis finished")
}
