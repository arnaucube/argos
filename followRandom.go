package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

/*func getUserByScreenName(ScreenName string) *twitter.User{


}*/
func getUserFollower(client *twitter.Client) string {
	ScreenName := "username"
	return ScreenName
}

func followUser(client *twitter.Client, ScreenName string) {

}

func followRandom(client *twitter.Client, nFollow int, ScreenName string) {
	fmt.Println("Starting to follow " + strconv.Itoa(nFollow) + " users")

	for i := 0; i < nFollow; i++ {
		ScreenName = getUserFollower(client)
		followUser(client, ScreenName)
	}

}

func optionFollowRandom(client *twitter.Client) {
	c.Red("how many accounts to follow?")
	newcommand := bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	fmt.Print("Number of users to follow: ")
	c.Purple(answer)
	nFollow, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("incorrect entry, need a positive number")
	}

	fmt.Print("entry @username of a user, to get a 1st user to follow, that will be a user that the 1st user is following, and the 2nd user will be a user that the 3rd user is following): @")
	newcommand = bufio.NewReader(os.Stdin)
	firstScreenName, _ := newcommand.ReadString('\n')
	firstScreenName = strings.TrimSpace(firstScreenName)
	fmt.Print("first user to follow: @")
	c.Purple(answer)

	c.Red("Are you sure? [y/n]")
	newcommand = bufio.NewReader(os.Stdin)
	answer, _ = newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		followRandom(client, nFollow, firstScreenName)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
