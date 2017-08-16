package main

type Example struct {
	Fieldname      string
	Anotherone     int
	AnotherExample *AnotherExample
}

type AnotherExample struct{}

func main() {
	ex := Example{}
	// Accessing a field
	ex.Fieldname
}
