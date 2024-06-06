package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-time.After(time.Second)
		}()
	}
	wg.Wait()
}
