package main

import (
	"fmt"
	"sync"
)

type threadsafeCounter struct {
	mu  sync.Mutex
	val int64
}

func createNewThreadsafeCounter() *threadsafeCounter {
	return &threadsafeCounter{}
}

func (c *threadsafeCounter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *threadsafeCounter) getVal() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}
func main() {
	threadSafeVal := createNewThreadsafeCounter()
	wg := sync.WaitGroup{}
	// Learning Always call wg.Add(N) before starting goroutines
	wg.Add(4)
	go func() {
		defer wg.Done()
		threadSafeVal.increment()
	}()
	go func() {
		defer wg.Done()
		threadSafeVal.increment()
	}()
	go func() {
		defer wg.Done()
		threadSafeVal.increment()
	}()
	go func() {
		defer wg.Done()
		threadSafeVal.increment()
	}()
	wg.Wait()
	fmt.Println("ThreadSafeVal value", threadSafeVal.getVal())
}
