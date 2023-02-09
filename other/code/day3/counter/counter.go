package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

const n = 10

func main() {
	var (
		count int64
		wg sync.WaitGroup
		mu sync.Mutex
	)
	
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				// atomic.AddInt64(&count, 1)

				mu.Lock()
				count++
				mu.Unlock()
				
			}
		}()
	}

	wg.Wait()

	fmt.Println(count)
}