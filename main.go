package main

import (
	"fmt"
	mutex "go_basic/cmd/14_mutex"
	"sync"
)

func main() {
	counter := mutex.NewCounter()

	var wg sync.WaitGroup

	for range 100 {
		wg.Go(func() {
			counter.Increment()
		})
	}

	wg.Wait()

	// 1. mutex 없이 하면, 레이스 컨디션 발생한다..
	// 	  go run -race .로 확인할수있음
	// 2. mutex 적용 후, 레이스 컨디션 발생 X
	fmt.Println(counter.Value())
}
