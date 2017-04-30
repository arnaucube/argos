package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func deleteTweets(client *twitter.Client, user *twitter.User) {
	tweets := getTweets(client, user.ScreenName, iterationsCount)
	count := 0
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
	c.Red("Are you sure you want to delete your tweets? [y/n]")
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
