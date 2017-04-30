package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const version = "0.1-dev"
const minNumWords = 10
const minNumHashtag = 2
const minNumUserInteractions = 1
const iterationsCount = 3

func main() {
	c.Yellow("Argos Panoptes")
	fmt.Println("---------------")
	c.Cyan("argos initialized")
	c.Purple("https://github.com/arnaucode/argos")
	fmt.Println("version " + version)
	fmt.Println("Reading twitterConfig.json file")
	client := readConfigTokensAndConnect()

	fmt.Println("---------------")
	fmt.Println("Getting user data...")
	user := getUserData(client)
	printUserData(user)
	fmt.Println("")

	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number")
	options := `
	1 - Analyze username
	2 - Unfollow all
	3 - Follow random
	4 - Delete Tweets
	5 - Delete Favs (Likes)
	6 - Tweet Random
	0 - Exit script
option to select: `
	for {
		fmt.Print(options)

		option, _ := newcommand.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("selected 1 - Analyze username")
			optionAnalyzeUserTweets(client)
			break
		case "2":
			fmt.Println("selected 2 - Unfollow all")
			optionUnfollowAll(client)
			break
		case "3":
			fmt.Println("selected 3 - Follow random")
			optionFollowRandom(client)
			break
		case "4":
			fmt.Println("selected 4 - Delete Tweets")
			optionDeleteTweets(client)
			break
		case "5":
			fmt.Println("selected 5 - Delete Favs (Likes)")
			optionDeleteFavs(client)
			break
		case "6":
			fmt.Println("selected 5 - Tweet random")
			optionTweetRandom(client)
			break
		case "0":
			fmt.Println("selected 0 - exit script")
			os.Exit(3)
			break
		default:
			fmt.Println("Invalid option")
			break
		}
	}
}
