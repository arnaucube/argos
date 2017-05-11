package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

var week = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func printBar(n int, lowest int, highest int) {
	bar := int((float64(n) / float64(highest)) * 40)

	if n == lowest {
		fmt.Print("\x1b[36;1m") //cyan
	}
	if n == highest {
		fmt.Print("\x1b[31;1m") //red
	}

	for i := 0; i < bar; i++ {
		fmt.Print("█")
	}

	if bar == 0 && n > 0 {
		fmt.Print("█")
	}
	if n == lowest {
		fmt.Print("	lowest")
	}
	if n == highest {
		fmt.Print("	highest")
	}
	fmt.Print("\x1b[0m") //defaultColor
	fmt.Println(" ")
}
func getHigherValueOfMap(m map[string]int) (int, int) {
	var values []int
	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)
	//returns lower, higher values
	return values[0], values[len(values)-1]
}

func printDays(days map[string]int) {
	lowest, highest := getHigherValueOfMap(days)

	for i := 0; i < len(week); i++ {
		fmt.Print(week[i] + " - " + strconv.Itoa(days[week[i]]) + "tw")
		fmt.Print("	")
		printBar(days[week[i]], lowest, highest)
	}
}
func analyzeDays(tweets []twitter.Tweet) {
	var days = make(map[string]int)
	for _, v := range tweets {
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
	lowest, highest := getHigherValueOfMap(hours)

	for i := 0; i < 24; i++ {
		var h string
		if i < 10 {
			h = "0" + strconv.Itoa(i)
		} else {
			h = strconv.Itoa(i)
		}
		fmt.Print(h + "h	-	" + strconv.Itoa(hours[h]) + "tw")
		fmt.Print("	")
		printBar(hours[h], lowest, highest)
	}
}
func analyzeHours(tweets []twitter.Tweet) {
	var hours = make(map[string]int)
	for _, v := range tweets {
		time := strings.Split(v.CreatedAt, " ")[3]
		hour := strings.Split(time, ":")[0]
		/*if _, ok := hours[hour]; ok {
			hours[hour] = hours[hour] + 1
		} else {
			hours[hour] = 1
		}*/
		hours[hour]++
	}
	printHours(hours)
}

func analyzeDates(tweets []twitter.Tweet) {
	fmt.Println("Weekly activity distribution (per day)")
	analyzeDays(tweets)
	fmt.Println(" ")
	fmt.Println("Daily activity distribution (per hour)")
	analyzeHours(tweets)
}
