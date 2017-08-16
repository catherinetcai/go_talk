package main

import "fmt"

func main() {
	x := map[string]string{"hello": "world"}

	for k, v := range x {
		fmt.Println(k, v)
	}
}
