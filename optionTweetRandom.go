package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func postTweet(client *twitter.Client, content string) {
	tweet, httpResp, err := client.Statuses.Update(content, nil)
	if err != nil {
		fmt.Println(err)
	}
	if httpResp.Status != "200 OK" {
		c.Red("error: " + httpResp.Status)
		c.Purple("maybe twitter has blocked the account, CTRL+C, wait 15 minutes and try again")
	}
	fmt.Print("tweet posted: ")
	c.Green(tweet.Text)
}
func tweetRandom(client *twitter.Client, nTweets int, screenName string) {
	fmt.Println("Starting to publish " + strconv.Itoa(nTweets) + " tweets")

	tweets := getTweets(client, screenName, 1)
	fmt.Println("the selected account have more than " + strconv.Itoa(len(tweets)))

	for i := 0; i < nTweets && i < len(tweets); i++ {
		postTweet(client, tweets[i].Text)
		waitTime(1)
	}

}

func optionTweetRandom(client *twitter.Client) {
	c.Red("ATTENTION!")
	c.Purple("Publishing tweets from a bot can be banned by twitter!")
	c.Cyan("Twitter can consider your account an spam account")
	c.Red("Maybe you can publish the random tweets by hand")
	c.Red("how many tweets to publish?")
	newcommand := bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	fmt.Print("Number of tweets to publish: ")
	c.Purple(answer)
	nTweets, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("incorrect entry, need a positive number")
	}

	fmt.Print("entry @username of a twitter account, the content of the tweets will be from the account tweets: @")
	newcommand = bufio.NewReader(os.Stdin)
	screenName, _ := newcommand.ReadString('\n')
	screenName = strings.TrimSpace(screenName)
	fmt.Print("user to get tweets content: @")
	c.Purple(screenName)

	c.Red("Are you sure? [y/n]")
	newcommand = bufio.NewReader(os.Stdin)
	answer, _ = newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		tweetRandom(client, nTweets, screenName)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
