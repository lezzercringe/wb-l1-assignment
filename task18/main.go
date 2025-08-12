package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32 = 0
	var wg sync.WaitGroup

	const timesPerWorker = 10000

	workersCount := runtime.NumCPU()
	for range workersCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range timesPerWorker {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}

	wg.Wait()
	expected := timesPerWorker * workersCount
	fmt.Printf("counter = %v (expected = %v)\n",
		atomic.LoadInt32(&counter),
		expected,
	)
}
