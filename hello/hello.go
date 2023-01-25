package main

import (
	"fmt"
	"log"

	"practice/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"sanath", "ramesh"}

	messages, err := greetings.Hellos(names)
	if err != nil {
        log.Fatal(err)
    }

	fmt.Println(messages)
}