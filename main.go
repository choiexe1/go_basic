package main

import (
	"fmt"
	functions "go_basic/cmd/03_functions"
)

func main() {
	fmt.Println(functions.Divide(10, 2))
	fmt.Println(functions.Sum(1, 2, 3, 4, 5))
	fmt.Println(functions.Sub(10, 5))
	fmt.Println(functions.Multiply(3, 4))
}
