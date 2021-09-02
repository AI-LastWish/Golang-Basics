package main

import "fmt"

type bot interface {
	// If there are any type in this program with a function called "getGreeting" and return a string, it is now an honorary member of type "bot"
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY CUSTOM logic
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY CUSTOM logic
	return "Hola!"
}
