package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

func optionAnalyzeUserFollowers(client *twitter.Client, username string) {
	var followers []twitter.User
	var maxid int64
	maxid = 0
	for i := 0; i < 4; i++ {
		followersRaw, _, err := client.Followers.List(&twitter.FollowerListParams{
			ScreenName: username,
			Count:      200,
			Cursor:     maxid,
		})
		check(err)
		maxid = followersRaw.NextCursor
		for _, f := range followersRaw.Users {
			/*fmt.Println("@" + follower.ScreenName)
			fmt.Println("CreatedAt: " + follower.CreatedAt)*/
			followers = append(followers, f)
		}
		fmt.Println(followersRaw.NextCursor)
		if followersRaw.NextCursor == 0 {
			break
		}
	}

	for k, follower := range followers {
		fmt.Print(k)
		fmt.Println("	@" + follower.ScreenName)
		fmt.Println("		CreatedAt: " + follower.CreatedAt)
	}
}
