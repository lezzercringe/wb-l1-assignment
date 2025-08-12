package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var workersCount int

func init() {
	flag.IntVar(&workersCount, "n", 5, "workers count")
}

func setupSignalHandler(handler func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-c
		handler()
	}()
}

func main() {
	flag.Parse()

	ch := make(chan int, workersCount)
	ctx, cancel := context.WithCancel(context.Background())

	setupSignalHandler(cancel)

	var wg sync.WaitGroup

	for i := range workersCount {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Shutting down worker %v\n", i)
					return
				case v := <-ch:
					fmt.Printf("Worker %v: %v\n", i, v)
				}
			}
		}()
	}

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		default:
			ch <- rand.Int()
		}
	}
}
