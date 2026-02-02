package main

import (
	"fmt"
	goroutines "go_basic/cmd/12_goroutines"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println(goroutines.DoA())
	}()

	go func() {
		defer wg.Done()
		fmt.Println(goroutines.DoB())
	}()

	go func() {
		defer wg.Done()
		fmt.Println(goroutines.DoC())
	}()

	wg.Wait()
}
