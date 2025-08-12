package main

import (
	"context"
	"runtime"
	"sync/atomic"
	"time"
)

func conditionalExit() {
	var done int32 = 0

	go func() {
		for atomic.LoadInt32(&done) == 0 {
			// do smth
		}
	}()

	<-time.After(2 * time.Second)
	atomic.StoreInt32(&done, 1)
}

func exitOnNotification(done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			// do smth
		}
	}
}

func exitOnContextDone(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// do something
		}
	}
}

func exitWithGoexit() {
	// do smth
	runtime.Goexit()
	// no more doing smth
}
