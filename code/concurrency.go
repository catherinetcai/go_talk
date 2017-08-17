package main

import "fmt"

func main() {
	go hello_world()
	fmt.Println("Main")
}

func hello_world() {
	fmt.Println("Hello world!")
}
