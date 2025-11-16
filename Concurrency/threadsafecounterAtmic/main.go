package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var a atomic.Int32
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		a.Add(1)
	}()
	go func() {
		defer wg.Done()
		a.Add(4)
	}()
	go func() {
		defer wg.Done()
		a.Add(5)
	}()
	go func() {
		defer wg.Done()
		a.Add(200)
	}()
	wg.Wait()
	fmt.Println("Existing Val", a.Load())

}
