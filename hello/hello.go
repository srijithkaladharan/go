package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")   //gets prefixed with each log message
	log.SetFlags(1)    // 0 - will not prefix the date, 1- will prefix the date

	names := []string{"Srijith","Mahith","Bheem"}

	//message, err := greetings.Hello("Srijith")
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)   //prints or consoles the error message and stops the program
	}


	// message := greetings.Hello("Srijith")
	fmt.Println(messages)
}