package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const version = "0.5-beta"
const minNumWords = 10
const minNumHashtag = 2
const minNumUserInteractions = 2
const iterationsCount = 3

func main() {
	eye := `
                ___________
            .-=d88888888888b=-.
        .:d8888pr |\|/-\| rq8888b.
      ,:d8888P^//\-\/_\ /_\/^q888/b.
    ,;d88888/~-/ .-~  _~-. |/-q88888b,
   //8888887-\ _/    (#)  \\-\/Y88888b\
   \8888888|// T           Y _/|888888 o
    \q88888|- \l           !\_/|88888p/
      q8888l\-//\         / /\|!8888P
        q888\/-|  -,___.-^\/-\/888P
          =88\./-/|/ |-/!\/-!/88=
            ^^ ------------- ^
           _____   _____  ____   _____
     /\   |  __ \ / ____|/ __ \ / ____|
    /  \  | |__) | |  __| |  | | (___
   / /\ \ |  _  /| | |_ | |  | |\___ \
  / ____ \| | \ \| |__| | |__| |____) |
 /_/    \_\_|  \_\\_____|\____/|_____/

 Open source twitter entropic toolkit
	`
	c.Cyan(eye)
	c.DarkGray("--Be half bot and half human, a new generation of cyborgs--")
	fmt.Println("---------------")
	fmt.Print("source code: ")
	c.Purple("https://github.com/arnaucode/argos")
	fmt.Print("project page: ")
	c.Purple("http://arnaucode/argos")
	fmt.Print("version ")
	c.Purple(version)
	fmt.Println("---------------")
	client := readConfigTokensAndConnect()

	fmt.Println("Getting user data...")
	user := getUserData(client)
	printUserData(user)
	if user.ScreenName == "" {
		c.Red("Can not connect to Twitter API, maybe the file twitterConfig.json is wrong")
		os.Exit(3)
	}
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
	7 - Analyze tweet
	8 - Analyze User Followers
	0 - Exit script
option to select: `
	for {
		fmt.Print(options)

		option, _ := newcommand.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("selected 1 - Analyze username")
			username := askForUsername()
			optionAnalyzeUserTweets(client, username)
			fmt.Println("")
			c.Purple("Note: the current hours displaying, are the Twitter servers hours (Coordinated Universal Time (UTC) +0000 UTC)")
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
			fmt.Println("selected 6 - Tweet random")
			optionTweetRandom(client)
			break
		case "7":
			fmt.Println("selected 7 - Analyze Tweet")
			optionAnalyzeTweet(client)
			break
		case "8":
			fmt.Println("selected 8 - Analyze User Followers")
			username := askForUsername()
			optionAnalyzeUserFollowers(client, username)
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
