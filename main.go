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
	fmt.Println("---------------")
	fmt.Print("\x1b[36;1m") //cyan
	fmt.Println("goTweetsAnalyze initialized")
	fmt.Print("\x1b[35;1m") //cyan
	fmt.Println("https://github.com/arnaucode/goTweetsAnalyze")
	fmt.Print("\x1b[0m") //defaultColor
	fmt.Println("version " + version)
	fmt.Println("Reading twitterConfig.json file")
	client := readConfigTokensAndConnect()
	fmt.Println("---------------")
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number")
	options := `
	1 - Analyze username
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
			fmt.Println("selected 2 - Delete all Favs")
			break
		case "3":
			fmt.Println("selected 3 - Delete all Tweets and Favs")
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
