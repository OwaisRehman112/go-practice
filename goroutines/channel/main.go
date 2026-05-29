package main

import "fmt"

func greet(ch chan string) {
	ch <- "Hello from goroutine"
}

func main() {

	// OPTION 1 (standard)
	ch := make(chan string)

	// OPTION 2 (buffered)
	// ch := make(chan string, 1)

	go greet(ch)

	msg := <-ch

	fmt.Println(msg)
}