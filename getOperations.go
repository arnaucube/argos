package main

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
)

func getTweets(client *twitter.Client, username string, iterations int) []twitter.Tweet {
	var tweets []twitter.Tweet
	var maxid int64
	for i := 0; i < iterations; i++ {
		tweetsRaw, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			ScreenName: username,
			Count:      200,
			MaxID:      maxid,
		})
		//if no tweets, stop getting tweets
		if len(tweetsRaw) == 0 {
			break
		}
		maxid = tweetsRaw[len(tweetsRaw)-1].ID

		for _, v := range tweetsRaw {
			tweets = append(tweets, v)
		}
	}

	return tweets
}

func getFavs(client *twitter.Client, username string, iterations int) []twitter.Tweet {
	var tweets []twitter.Tweet
	var maxid int64
	for i := 0; i < iterations; i++ {
		tweetsRaw, _, _ := client.Favorites.List(&twitter.FavoriteListParams{
			ScreenName: username,
			Count:      200,
			MaxID:      maxid,
		})

		//if no tweets, stop getting tweets
		if len(tweetsRaw) == 0 {
			break
		}
		maxid = tweetsRaw[len(tweetsRaw)-1].ID

		for _, v := range tweetsRaw {
			tweets = append(tweets, v)
		}
	}

	return tweets
}

func getUserData(client *twitter.Client) *twitter.User {
	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	return user
}

func printUserData(user *twitter.User) {
	fmt.Print("username: ")
	c.Cyan(user.Name + " @" + user.ScreenName)
	if user.Email != "" {
		fmt.Print("Email ")
		c.Red(user.Email)
	}
	if user.Location != "" {
		fmt.Print("Location: ")
		c.Red(user.Location)
	}
	fmt.Print("user created on: ")
	c.Cyan(user.CreatedAt)

	fmt.Print("number of tweets: ")
	c.Purple(strconv.Itoa(user.StatusesCount))

	fmt.Print("number of favs: ")
	c.Purple(strconv.Itoa(user.FavouritesCount))
}

func getRetweets(client *twitter.Client, tweetId int64) []twitter.Tweet {
	var tweets []twitter.Tweet
	tweets, _, err := client.Statuses.Retweets(tweetId, &twitter.StatusRetweetsParams{
		Count: 200,
	})
	if err != nil {
		fmt.Println(err)
	}
	return tweets
}
