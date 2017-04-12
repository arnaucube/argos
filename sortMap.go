package main

import (
	"fmt"
	"sort"
)

func sortMap(m map[string]int) map[int][]string {
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		if v > 2 {
			n[v] = append(n[v], k)
		}
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%s, %d\n", s, k)
		}
	}
	return n
}
