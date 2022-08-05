package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)


func main() {
	message, err := greetings.Hello("Pero")

	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println(message)
}