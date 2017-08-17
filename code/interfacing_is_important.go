package main

import "fmt"

type CloudProvider interface {
	Type() string
}

type AWS struct{}

func (a *AWS) Type() string { return "aws" }

type GCE struct{}

func (g *GCE) Type() string { return "gce" }

func PrintType(c CloudProvider) {
	fmt.Println(c.Type())
}
