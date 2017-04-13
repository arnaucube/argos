package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

var week = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func printBar(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("█")
	}
	fmt.Println(" ")
}

func printDays(days map[string]int) {
	for i := 0; i < len(week); i++ {
		fmt.Print(week[i] + " - " + strconv.Itoa(days[week[i]]) + "tw")
		fmt.Print("	")
		printBar(days[week[i]])
	}
}
func analyzeDays(tweets []twitter.Tweet) {
	var days = make(map[string]int)
	for _, v := range tweets {
		//fmt.Println(v.CreatedAt)
		day := strings.Split(v.CreatedAt, " ")[0]
		if _, ok := days[day]; ok {
			days[day] = days[day] + 1
		} else {
			days[day] = 1
		}
	}
	printDays(days)
}

func printHours(hours map[string]int) {
	for i := 0; i < 24; i++ {
		var h string
		if i < 10 {
			h = "0" + strconv.Itoa(i)
		} else {
			h = strconv.Itoa(i)
		}
		fmt.Print(h + "h	-	" + strconv.Itoa(hours[h]) + "tw")
		fmt.Print("	")
		printBar(hours[h])
	}
}
func analyzeHours(tweets []twitter.Tweet) {
	var hours = make(map[string]int)
	for _, v := range tweets {
		time := strings.Split(v.CreatedAt, " ")[3]
		hour := strings.Split(time, ":")[0]
		if _, ok := hours[hour]; ok {
			hours[hour] = hours[hour] + 1
		} else {
			hours[hour] = 1
		}
	}
	printHours(hours)
}

func analyzeDates(tweets []twitter.Tweet) {
	analyzeDays(tweets)
	fmt.Println(" ")
	analyzeHours(tweets)
}
