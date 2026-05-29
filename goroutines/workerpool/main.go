package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	time.Sleep(1 * time.Second)

	ch <- fmt.Sprintf("Worker %d done", id)
}

func main() {

	// OPTION 1 (recommended)
	var wg sync.WaitGroup

	// OPTION 2
	// wg := &sync.WaitGroup{}

	ch := make(chan string)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg, ch)
	}

	// closing channel after workers finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("All complete")
}