package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Config struct {
	Consumer_key        string `json:"consumer_key"`
	Consumer_secret     string `json:"consumer_secret"`
	Access_token_key    string `json:"access_token_key"`
	Access_token_secret string `json:"access_token_secret"`
}

func readConfigTokensAndConnect() (client *twitter.Client) {
	var config Config
	file, e := ioutil.ReadFile("twitterConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &config)
	//fmt.Printf("%+v\n", config)
	fmt.Println("twitterConfig.json read comlete")

	fmt.Print("connecting to twitter api --> ")
	configu := oauth1.NewConfig(config.Consumer_key, config.Consumer_secret)
	token := oauth1.NewToken(config.Access_token_key, config.Access_token_secret)
	httpClient := configu.Client(oauth1.NoContext, token)
	// twitter client
	client = twitter.NewClient(httpClient)

	fmt.Println("connection successfull")

	return client
}

func analyzeUsername(client *twitter.Client) {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("enter username: @")
	username, _ := newcommand.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Println("user selected: " + username)

	// Home Timeline
	user, resp, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: username,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("%+v\n", user)
	fmt.Println(resp)
}

func main() {
	fmt.Println("---------------")
	fmt.Println("Tweets and favs delete script. Starting.")
	fmt.Println("Reading twitterConfig.json file")
	client := readConfigTokensAndConnect()

	fmt.Println("---------------")
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number:")
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
			analyzeUsername(client)
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
