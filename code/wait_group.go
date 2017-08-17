package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello_world(&wg)
	wg.Wait()
	fmt.Println("Main")
}

func hello_world(wg *sync.WaitGroup) {
	fmt.Println("Hello world!")
	wg.Done()
}

// Output:
// Hello world!
// Main
