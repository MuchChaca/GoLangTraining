package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns a map of the counts of each “word” in the string s
func WordCount(s string) map[string]int {
	strSlice := strings.Fields(s)
	strMap := make(map[string]int)
	for _, str := range strSlice {
		if _, ok := strMap[str]; ok {
			strMap[str]++
		} else {
			strMap[str] = 1
		}
	}
	return strMap
}

func main() {
	wc.Test(WordCount)
}
