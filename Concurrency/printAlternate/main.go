package main

import (
	"fmt"
	"sync"
)

const RANGE = 20

func printOdd(evenChan chan struct{}, oddChan chan struct{}) {
	for i := 1; i < RANGE; i += 2 {
		<-evenChan
		fmt.Println(i)
		oddChan <- struct{}{}
	}
}

func printEven(evenChan chan struct{}, oddChan chan struct{}) {
	for i := 0; i < RANGE; i += 2 {
		<-oddChan
		fmt.Println(i)
		evenChan <- struct{}{}
	}
}

func main() {
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		oddChan <- struct{}{}
		printOdd(evenChan, oddChan)
	}()
	go func() {
		defer wg.Done()
		printEven(evenChan, oddChan)
		<-oddChan
	}()
	wg.Wait()
}
