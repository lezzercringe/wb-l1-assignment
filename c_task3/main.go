package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var workersCount int

func init() {
	flag.IntVar(&workersCount, "n", 5, "workers count")
}

func main() {
	flag.Parse()

	ch := make(chan int, workersCount)

	for i := range workersCount {
		go func() {
			for v := range ch {
				fmt.Printf("Worker %v: %v\n", i, v)
			}
		}()
	}

	for {
		ch <- rand.Int()
	}
}
