package main

import (
	"fmt"
	"log"
	"example.com/greetings"
	
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Glandys", "Samantha", "Edmilson"}

	message, err := greetings.Hellos(names)

	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println(message)
}