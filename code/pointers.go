package main

import "fmt"

func main() {
	a := 1
	i := &a        // i is the value of the pointer of a
	*i = 21        // Sets what i points to to the value of 21
	fmt.Println(a) // What will this print?
}
