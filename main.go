package main

import (
	"fmt"
	librarysystem "go_basic/cmd/07_structs"
)

func main() {
	book := librarysystem.Book{
		Title:  "Go Programming Book",
		Author: "Rob Pike, Ken Thopson, Robert Griesemer",
		ISBN:   "1",
	}

	member := librarysystem.Member{
		ID:   "1",
		Name: "Jay Choi",
	}

	fmt.Println(librarysystem.NewLoan(member, book, "2026-02-03"))
}
