package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Ayoub")
	fmt.Println(message)

}
