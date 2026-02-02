package main

import (
	"fmt"
	goroutines "go_basic/cmd/12_goroutines"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://asdfzx87cv98a7d",
		"https://asdf7as9d8fas8df",
	}

	results := goroutines.FetchAll(urls)
	for _, r := range results {
		fmt.Println(r)
	}
}
