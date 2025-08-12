package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func writer(ctx context.Context, ch chan<- int) {
	defer close(ch)

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(200 * time.Millisecond): // rate limiting
			ch <- rand.Int()
		}
	}
}

func main() {
	const timeout = 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan int)

	go writer(ctx, ch)

	for {
		select {
		case <-ctx.Done():
			return
		case v := <-ch:
			fmt.Println(v)
		}
	}
}
