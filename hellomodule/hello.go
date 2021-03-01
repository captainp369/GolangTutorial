package main

import (
	"fmt"

	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Satah", "Darril"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
