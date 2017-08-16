package main

import (
	"errors"
	"log"
)

func AnErrorFunction() (string, error) {
	return "", errors.New("Test error")
}

func main() {
	b, err := AnErrorFunction()
	if err != nil {
		// This will log the error to standard out, and then exit with an error
		log.Fatal(err)
	}
}
