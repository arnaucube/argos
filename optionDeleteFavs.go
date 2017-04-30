package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func deleteFavs(client *twitter.Client, user *twitter.User) {
	tweets := getFavs(client, user.ScreenName, iterationsCount)
	count := 0
	for _, v := range tweets {
		c.Red("deleting: [id: " + v.IDStr + "] " + v.Text)
		deleted, _, _ := client.Favorites.Destroy(&twitter.FavoriteDestroyParams{
			ID: v.ID,
		})
		c.Green("DELETED: [id: " + deleted.IDStr + "] " + deleted.Text)
		count++
	}
	c.Red("Deleted " + strconv.Itoa(count) + " favs")
}

func optionDeleteFavs(client *twitter.Client) {
	c.Red("Are you sure you want to delete your favs? [y/n]")
	newcommand := bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		user := getUserData(client)
		deleteFavs(client, user)
		user = getUserData(client)
		printUserData(user)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
