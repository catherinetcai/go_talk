package main

import "fmt"

type Square struct {
	Length int
	Width  int
}

func (s Square) Area() int {
	return s.Length * s.Width
}

func main() {
	s := Square{2, 2}
	fmt.Println(s.Area())
}
