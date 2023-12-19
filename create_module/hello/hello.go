package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("log: ")
	log.SetFlags(0)


	// greetingMessage, err := greetings.Hello("")
	names := []string{"Hiru", "John", "Doe"}

	greetingMessages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(greetingMessages)
}
