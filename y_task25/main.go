package main

import (
	"fmt"
	"runtime"
	"time"
)

func sleepSimple(d time.Duration) {
	<-time.After(d)
}

func sleepBarbarian(d time.Duration) {
	end := time.Now().Add(d)
	for time.Now().Before(end) {
		runtime.Gosched() // чтобы не настолько по-варварски
	}
}

func main() {
	start := time.Now()
	sleepSimple(3 * time.Second)
	fmt.Println(time.Since(start))

	start = time.Now()
	sleepBarbarian(3 * time.Second)
	fmt.Println(time.Since(start))
}
