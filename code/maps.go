package main

import "fmt"

func main() {
	x := map[string]string{"hello": "world"}

	for k, v := range x {
		fmt.Println(k, v)
	}

	// Ok will tell you if the value exists
	y, ok := x["hello"]
}
