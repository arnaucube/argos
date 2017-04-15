package main

import (
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func analyzeSource(tweets []twitter.Tweet) map[string]int {
	var sources = make(map[string]int)
	for _, v := range tweets {
		source := strings.Split(strings.Split(v.Source, ">")[1], "<")[0]
		if _, ok := sources[source]; ok {
			sources[source] = sources[source] + 1
		} else {
			sources[source] = 1
		}
	}
	return (sources)
}
