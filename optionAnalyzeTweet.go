package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

type DateRT struct {
	Retweets []twitter.Tweet
}

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

	tweet, _, err := client.Statuses.Show(tweetId, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("tweet text: ")
	c.Yellow(tweet.Text)

	retweets := getRetweets(client, tweetId)

	var dates = make(map[string]DateRT)
	for _, retweet := range retweets {
		retws := dates[retweet.CreatedAt].Retweets
		retws = append(retws, retweet)
		var currDate DateRT
		currDate.Retweets = retws
		dates[retweet.CreatedAt] = currDate
	}
	fmt.Print("total of: ")
	fmt.Println(len(retweets))

	for k, v := range dates {
		printDateRT(k, v)
	}
	c.Purple("Warning: Twitter API only gives the last 100 Retweets")
}
func printDateRT(k string, v DateRT) {
	fmt.Print("\x1b[32;1m" + k + "\x1b[0m")
	fmt.Println("	" + strconv.Itoa(len(v.Retweets)) + " Retweets at this date:")
	for _, retweet := range v.Retweets {
		source := strings.Split(strings.Split(retweet.Source, ">")[1], "<")[0]
		if len(v.Retweets) > 1 {
			fmt.Print("\x1b[31;1m") //red
		}
		fmt.Print("		@" + retweet.User.ScreenName)
		fmt.Print("\x1b[0m") //defaultColor
		fmt.Println(",	source: " + source)

	}

}
