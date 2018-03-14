package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

type DateRT struct {
	Retweets []twitter.Tweet
	Date     string
}

func printDateRT(v DateRT) {
	fmt.Print("\x1b[32;1m" + v.Date + "\x1b[0m")
	fmt.Println("	" + strconv.Itoa(len(v.Retweets)) + " Retweets at this date:")
	for _, retweet := range v.Retweets {
		source := strings.Split(strings.Split(retweet.Source, ">")[1], "<")[0]
		if len(v.Retweets) > 1 {
			fmt.Print("\x1b[31;1m") //red
		}
		fmt.Print("	@" + retweet.User.ScreenName)
		fmt.Print("\x1b[0m") //defaultColor
		fmt.Print("	(ID: ")
		fmt.Print("\x1b[31;1m" + retweet.User.IDStr + "\x1b[0m)")
		if source == "TweetDeck" {
			fmt.Print(",	source: \x1b[33;1m" + source + "\x1b[0m")
		} else {
			fmt.Print(",	source: " + source)
		}
		fmt.Print(",	user created at: \x1b[32;1m" + retweet.User.CreatedAt + "\x1b[0m,")

		fmt.Print("	\x1b[34;1m" + strconv.Itoa(retweet.User.FollowersCount) + "\x1b[0m followers")
		fmt.Print(", \x1b[34;1m" + strconv.Itoa(retweet.User.FriendsCount) + "\x1b[0m following")
		fmt.Println("")
	}
	fmt.Println("")

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
	fmt.Print("tweet date: ")
	c.Green(tweet.CreatedAt)

	retweets := getRetweets(client, tweetId)

	var dates = make(map[string]DateRT)
	for _, retweet := range retweets {
		createdat, err := retweet.CreatedAtTime()
		check(err)
		createdatRounded := time.Date(createdat.Year(), createdat.Month(), createdat.Day(), createdat.Hour(), createdat.Minute(), 0, 0, createdat.Location())

		retws := dates[createdatRounded.String()].Retweets
		retws = append(retws, retweet)
		var currDate DateRT
		currDate.Retweets = retws
		currDate.Date = createdatRounded.String()
		dates[createdatRounded.String()] = currDate
	}
	fmt.Println("total of " + strconv.Itoa(len(retweets)) + " retweets")

	//puts the map into a slice, to easy sort
	var arrayDates []DateRT
	for _, v := range dates {
		arrayDates = append(arrayDates, v)
	}

	fmt.Println("")
	c.Cyan("Showing accounts that retweeted at the same exact time (they are possible Bots) and the source the RT were made:")

	//sort the slice
	sort.Slice(arrayDates, func(i, j int) bool {
		return len(arrayDates[i].Retweets) < len(arrayDates[j].Retweets)
	})
	//print the sorted slice
	for _, v := range arrayDates {
		if len(v.Retweets) > 1 {
			printDateRT(v)
		}
	}
	fmt.Println("")
	c.Purple("Warning: Twitter API only gives the last 100 Retweets")
}
