package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	//get tweets and analyze words and dates
	tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: username,
		Count:      200,
	})

	fmt.Println("word count")
	analyzeWords(tweets)

	analyzeDates(tweets)
	fmt.Println("")
	fmt.Println("Devices:")
	sources := analyzeSource(tweets)
	for k, v := range sources {
		fmt.Print(k + ": ")
		fmt.Println(strconv.Itoa(v) + "tw	")
	}

	fmt.Println(" ")
	fmt.Print("first tweet analyzed: ")
	fmt.Println(tweets[len(tweets)-1].CreatedAt)
	fmt.Print("last tweet analyzed: ")
	fmt.Println(tweets[0].CreatedAt)

	fmt.Println(" ")
	fmt.Println("User @" + username + " analysis finished")
}
