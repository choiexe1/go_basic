package main

import (
	"fmt"
	pointerprac "go_basic/cmd/08_pointers"
)

func main() {
	l := pointerprac.NewLinkedList()

	l.Insert(1)
	l.Insert(4)
	l.Insert(9)
	l.Insert(7)

	fmt.Println(l.Find(7))
	fmt.Println(l.Delete(7))
	fmt.Println(l.Find(7))
	fmt.Println(l)
}
