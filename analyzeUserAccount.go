package main

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
)

func analyzeUserAccount(client *twitter.Client, username string) {
	user, _, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: username,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("User created at: ")
	c.Cyan(user.CreatedAt)

	fmt.Print("email: ")
	c.Cyan(user.Email)

	fmt.Print("Following count: ")
	c.Cyan(strconv.Itoa(user.FriendsCount))
	fmt.Print("Followers count: ")
	c.Cyan(strconv.Itoa(user.FollowersCount))

	fmt.Print("Location: ")
	c.Cyan(user.Location)

	fmt.Print("Timezone: ")
	c.Cyan(user.Timezone)
	fmt.Println("")
	fmt.Println("")
}
