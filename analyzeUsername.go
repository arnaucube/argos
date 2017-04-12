package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

type Tweet struct {
	text string `json: "Text"`
}

var words = make(map[string]int)

func mapWords(text string) {
	s := strings.Split(text, " ")

	for _, v := range s {
		if _, ok := words[v]; ok {
			words[v] = words[v] + 1
		} else {
			words[v] = 1
		}
	}
}

func analyzeUsername(client *twitter.Client) {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("enter username: @")
	username, _ := newcommand.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Println("user selected: " + username)

	fmt.Println("-----------------------")

	tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: username,
		Count:      200,
	})

	for _, v := range tweets {
		mapWords(v.Text)
	}

	fmt.Println(len(words))
	_ = sortMap(words)
}
