package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	w := strings.Fields(s)
	wc := make(map[string]int)
	for _, v := range w {
		wc[v]++
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
