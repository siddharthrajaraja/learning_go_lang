package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}

type SpanishBot struct{}

func (EnglishBot) getGreeting() string {
	// Some logic for Eng
	return "Hello world"
}

func (SpanishBot) getGreeting() string {
	// Some logic for Spanish
	return "Hola world"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
