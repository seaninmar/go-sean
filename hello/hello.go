package main

import (
	"fmt"
	"github.com/seaninmar/go-sean/greetings"
	"log"
	"os"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{os.Getenv("USER"), "Gladys", "Samantha", "Darrin"}

	// Request a greeting message.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	fmt.Println("-- Quotes --")
	fmt.Println("rsc.io/quote.Go()\t", quote.Go())
	fmt.Println("rsc.io/quote.Glass()\t", quote.Glass())
	fmt.Println("rsc.io/quote.Hello()\t", quote.Hello())
	fmt.Println("rsc.io/quote.Opt()\t", quote.Opt())
}
