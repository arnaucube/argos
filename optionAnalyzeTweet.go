package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func optionAnalyzeTweet(client *twitter.Client) {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("enter link of the tweet: ")
	link, _ := newcommand.ReadString('\n')
	link = strings.TrimSpace(link)
	fmt.Print("link selected: ")
	c.Cyan(link)
	fmt.Println("-----------------------")
	linkParam := strings.Split(link, "/")
	tweetIdStr := linkParam[len(linkParam)-1]
	c.Cyan(tweetIdStr)
	tweetId, err := strconv.ParseInt(tweetIdStr, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	tweets := getRetweets(client, tweetId)
	for _, tweet := range tweets {
		source := strings.Split(strings.Split(tweet.Source, ">")[1], "<")[0]
		fmt.Println(tweet.CreatedAt + " @" + tweet.User.ScreenName + ", source: " + source)
	}
	fmt.Print("total of: ")
	fmt.Println(len(tweets))
}
