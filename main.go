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
const iterationsCount = 3

func main() {
	c.DarkGray("Argus Panoptes")
	fmt.Println("---------------")
	c.Cyan("argus initialized")
	c.Purple("https://github.com/arnaucode/argus")
	fmt.Println("version " + version)
	fmt.Println("Reading twitterConfig.json file")
	client := readConfigTokensAndConnect()
	fmt.Println("---------------")
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number")
	options := `
	1 - Analyze username
	2 - Delete Tweets
	0 - Exit script
option to select: `
	for {
		fmt.Print(options)

		option, _ := newcommand.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("selected 1 - Analyze username")
			getUserTweets(client)
			break
		case "2":
			fmt.Println("selected 2 - Delete Tweets")
			deleteTweetsAndFavs(client)
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
