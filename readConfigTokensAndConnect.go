package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//Config stores the data from json twitterConfig.json file
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
