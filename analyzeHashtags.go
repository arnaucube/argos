package main

import (
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func analyzeHashtags(tweets []twitter.Tweet, words map[string]int) map[string]int {
	var hashtags = make(map[string]int)
	for v, k := range words {
		if strings.Contains(v, "#") {
			hashtags[v] = k
		}
	}
	return (hashtags)
}
