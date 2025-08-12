package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 10}
	var wg sync.WaitGroup
	for _, num := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(num * num)
		}()
	}

	wg.Wait()
}
