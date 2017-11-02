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
func unfollowUser(client *twitter.Client, screenName string) {
	user, httpResp, _ := client.Friendships.Destroy(&twitter.FriendshipDestroyParams{
		ScreenName: screenName,
	})
	if httpResp.Status != "200 OK" {
		c.Red(httpResp.Status)
	}
	c.Green(user.ScreenName)
}
func getFollowingUsers(client *twitter.Client, user *twitter.User) {
	following, httpResp, err := client.Friends.List(&twitter.FriendListParams{
		ScreenName: user.ScreenName,
		Count:      200,
	})
	if err != nil {
		fmt.Println(err)
	}
	if httpResp.Status != "200 OK" {
		c.Red(httpResp.Status)
	}
	for _, k := range following.Users {
		unfollowUser(client, k.ScreenName)
		//wait to avoid the twitter api limitation
		/*fmt.Println("waiting 1 min to avoid twitter api limitation")
		fmt.Println(time.Now().Local())
		time.Sleep(1 * time.Minute)*/
	}
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
		getFollowingUsers(client, user)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
