package main

// interface{} is a type to hold any value
var foo interface{} = "string"

func main() {
	// You have to use type assertion in order to "convert" interface{} to a type

	// This is ok
	foo.(string)

	// This will panic, because
	foo.(int)

	switch foo.(type) {
	case string:
		// stuff
	case int:
		// more stuff
	default:
	}
}
