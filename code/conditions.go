package main

import "fmt"

func main() {
	if 1 < 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// Switch statements are also a thing
	switch foo {
	case 'a':
	case 'b':
	default:
	}
}
