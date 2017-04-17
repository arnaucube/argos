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
		//if no tweets, stop getting tweets
		if len(tweetsRaw) == 0 {
			break
		}
		maxid = tweetsRaw[len(tweetsRaw)-1].ID

		for _, v := range tweetsRaw {
			tweets = append(tweets, v)
		}
	}

	return tweets
}
func getUserTweets(client *twitter.Client) {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("enter username: @")
	username, _ := newcommand.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Print("user selected: ")
	c.Cyan("@" + username)
	fmt.Println("-----------------------")

	//get tweets
	tweets := getTweets(client, username, iterationsCount)

	if len(tweets) == 0 {
		fmt.Println("User @" + username + " does not have tweets")
		return
	}
	//now analyze words and dates
	fmt.Println("Word frequency (more than " + strconv.Itoa(minNumWords) + " times):")
	words := analyzeWords(tweets)

	fmt.Println("Hashtags used (more than " + strconv.Itoa(minNumHashtag) + " times): ")
	hashtags := analyzeHashtags(tweets, words)
	printSortedMapStringInt(hashtags, minNumHashtag)
	fmt.Println("")

	analyzeDates(tweets)
	fmt.Println("")
	fmt.Println("Devices:")
	sources := analyzeSource(tweets)
	for k, v := range sources {
		fmt.Print("\x1b[32;1m") //cyan
		fmt.Print(k + ": ")
		fmt.Print("\x1b[0m") //defaultColor
		fmt.Println(strconv.Itoa(v) + "tw	")
	}

	fmt.Println(" ")
	fmt.Print("first tweet analyzed: ")
	fmt.Println(tweets[len(tweets)-1].CreatedAt)
	fmt.Print("last tweet analyzed: ")
	fmt.Println(tweets[0].CreatedAt)

	fmt.Println(" ")
	fmt.Println("Total of " + strconv.Itoa(len(tweets)) + " tweets analyzed")
	fmt.Println(" ")
	fmt.Print("User @")
	c.Cyan(username)
	fmt.Println(" analysis finished")
}
