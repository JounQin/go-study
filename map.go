package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	fmt.Println(words)
	for _, word := range words {
	   result[word] += 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
