package main

import (
	"fmt"
	channels "go_basic/cmd/13_channels"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://asdfzx87cv98a7d",
		"https://asdf7as9d8fas8df",
	}

	result := channels.FetchOne(urls[0])
	fmt.Println(result)
}
