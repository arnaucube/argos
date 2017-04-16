package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func getTweets(client *twitter.Client, username string, iterations int) []twitter.Tweet {
	var tweets []twitter.Tweet
	var maxid int64
	for i := 0; i < iterations; i++ {
		tweetsRaw, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			ScreenName: username,
			Count:      200,
			MaxID:      maxid,
		})
		maxid = tweetsRaw[len(tweetsRaw)-1].ID

		for _, v := range tweetsRaw {
			tweets = append(tweets, v)
		}
	}

	fmt.Println("total of " + strconv.Itoa(len(tweets)) + " tweets")
	return tweets
}
func getUser(client *twitter.Client) {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("enter username: @")
	username, _ := newcommand.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Println("user selected: " + username)

	fmt.Println("-----------------------")

	//get tweets
	tweets := getTweets(client, username, 2)

	//now analyze words and dates
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
