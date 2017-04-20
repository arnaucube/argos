package main

import (
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func analyzeUserInteractions(tweets []twitter.Tweet, words map[string]int) map[string]int {
	var userInteractions = make(map[string]int)
	for v, k := range words {
		if strings.Contains(v, "@") {
			userInteractions[v] = k
		}
	}
	return (userInteractions)
}
