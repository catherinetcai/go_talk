package main

type Thing struct {
	FieldA string
	FieldB string
	FieldC string
}

func main() {
	thing := &Thing{
		FieldA: "foo",
		FieldB: "bar",
		FieldC: "baz",
	}
}
