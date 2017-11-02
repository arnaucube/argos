package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func askForUsername() string {
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("enter username: @")
	username, _ := newcommand.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Print("user selected: ")
	c.Cyan("@" + username)
	fmt.Println("-----------------------")
	return username
}
