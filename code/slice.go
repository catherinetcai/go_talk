package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	for i := range x {
		fmt.Println(i)
	}
	// Or
	y := make([]int, 5) // Allocates an empty array
}
