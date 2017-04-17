package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func getAndPrintUserData(client *twitter.Client) *twitter.User {
	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
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
	c.Cyan("user created on: " + user.CreatedAt)

	fmt.Print("number of tweets: ")
	c.Purple(strconv.Itoa(user.StatusesCount))
	return user
}
func deleteTweets(client *twitter.Client, user *twitter.User) {
	tweets := getTweets(client, user.ScreenName, iterationsCount)
	for _, v := range tweets {
		c.Red("deleting: [id: " + v.IDStr + "] " + v.Text)
		deleted, _, _ := client.Statuses.Destroy(v.ID, nil)
		c.Green("deleting: [id: " + deleted.IDStr + "] " + deleted.Text)
	}
}
func deleteFavs(client *twitter.Client) {

}
func deleteTweetsAndFavs(client *twitter.Client) {
	fmt.Println("Getting user data...")
	user := getAndPrintUserData(client)
	fmt.Println("")
	fmt.Println("Are you sure you want to delete you tweets? [y/n]")
	newcommand := bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		deleteTweets(client, user)
		deleteFavs(client)
		user = getAndPrintUserData(client)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
