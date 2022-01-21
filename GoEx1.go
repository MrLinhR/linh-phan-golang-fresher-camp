package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int, 1001)
var wg = sync.WaitGroup{}

func main() {

	for i := 1; i <= 1000; i++ {
		ch <- i
	}
	wg.Add(5)
	go crawlData(1)
	go crawlData(2)
	go crawlData(3)
	go crawlData(4)
	go crawlData(5)
	close(ch)
	wg.Wait()
}

func crawlData(number int) {
	for data := range ch {
		fmt.Println("Goroutine ", number, ": ", data)
	}
	wg.Done()
}
