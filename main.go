package main

import (
	"fmt"
	stack "go_basic/cmd/05_slices"
)

func main() {
	stack.Push(1)
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
}
