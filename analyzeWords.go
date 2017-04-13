package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func sortMapWords(m map[string]int) map[int][]string {
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		if v > minNumWords {
			n[v] = append(n[v], k)
		}
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%d - %s,\n", k, s)
		}
	}
	return n
}

func mapWords(text string, words map[string]int) {
	s := strings.Split(text, " ")

	for _, v := range s {
		if _, ok := words[v]; ok {
			words[v] = words[v] + 1
		} else {
			words[v] = 1
		}
	}
}
func analyzeWords(tweets []twitter.Tweet) {
	var words = make(map[string]int)

	for _, v := range tweets {
		mapWords(v.Text, words)
	}

	fmt.Println(len(words))
	//get sorted list of frequency words
	_ = sortMapWords(words)
	fmt.Println(" ")
}
