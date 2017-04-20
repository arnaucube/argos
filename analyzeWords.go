package main

import (
	"fmt"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func mapWords(text string, words map[string]int) map[string]int {
	s := strings.Split(text, " ")

	for _, v := range s {
		if _, ok := words[v]; ok {
			words[v] = words[v] + 1
		} else {
			words[v] = 1
		}
	}
	return words
}
func analyzeWords(tweets []twitter.Tweet) map[string]int {
	var words = make(map[string]int)

	for _, v := range tweets {
		words = mapWords(v.Text, words)
	}

	//get sorted list of frequency words
	_ = printSortedMapStringInt(words, minNumWords)
	fmt.Println(" ")
	return words
}
