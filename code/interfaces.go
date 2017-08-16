package main

import "fmt"

type Foo interface {
	Foo()
}

type Bar struct{}

func (b *Bar) Foo() {
	fmt.Println("Foo!")
}

type AnotherBar struct{}

func Baz(f Foo) {}

func main() {
	b := &Bar{}
	Baz(b)
}
