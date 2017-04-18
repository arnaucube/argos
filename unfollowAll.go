package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func printUserFollowsData(user *twitter.User) {
	fmt.Print("followers: ")
	c.Cyan(strconv.Itoa(user.FollowersCount))
	fmt.Print("following: ")
	c.Cyan(strconv.Itoa(user.FriendsCount))
}
func unfollowFollowingUsers(client *twitter.Client, user *twitter.User) {
	following, _, _ := client.Friends.List(&twitter.FriendListParams{
		ScreenName: user.ScreenName,
		Count:      200,
	})
	fmt.Println(following)
}
func optionUnfollowAll(client *twitter.Client) {
	fmt.Println("Getting user data...")
	user := getUserData(client)
	printUserFollowsData(user)
	fmt.Println("")
	c.Red("Are you sure you want to unfollow all? [y/n]")
	newcommand := bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		unfollowFollowingUsers(client, user)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
