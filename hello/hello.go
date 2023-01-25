package main

import (
	"fmt"
	"practice/greetings"
)

func main() {
	message := greetings.Hello("sanath")
	fmt.Println(message)
}