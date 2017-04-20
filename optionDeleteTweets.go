package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func getUserData(client *twitter.Client) *twitter.User {
	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	return user
}
func printUserData(user *twitter.User) {
	fmt.Print("username: ")
	c.Cyan(user.Name + " @" + user.ScreenName)
	if user.Email != "" {
		fmt.Print("Email ")
		c.Red(user.Email)
	}
	if user.Location != "" {
		fmt.Print("Location: ")
		c.Red(user.Location)
	}
	fmt.Print("user created on: ")
	c.Cyan(user.CreatedAt)

	fmt.Print("number of tweets: ")
	c.Purple(strconv.Itoa(user.StatusesCount))
}
func deleteTweets(client *twitter.Client, user *twitter.User) {
	tweets := getTweets(client, user.ScreenName, iterationsCount)
	count := 0
	//fmt.Println(tweets)
	for _, v := range tweets {
		c.Red("deleting: [id: " + v.IDStr + "] " + v.Text)
		deleted, _, _ := client.Statuses.Destroy(v.ID, nil)
		c.Green("DELETED: [id: " + deleted.IDStr + "] " + deleted.Text)
		count++
	}
	c.Red("Deleted " + strconv.Itoa(count) + " tweets")
}
func optionDeleteTweets(client *twitter.Client) {
	fmt.Println("Getting user data...")
	user := getUserData(client)
	printUserData(user)
	fmt.Println("")
	c.Red("Are you sure you want to delete you tweets? [y/n]")
	newcommand := bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		deleteTweets(client, user)
		user = getUserData(client)
		printUserData(user)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
