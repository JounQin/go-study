package main

import (
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(messages)
}
