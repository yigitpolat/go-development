package main

import "fmt"

type bot interface {
	getGreeting() string
	// getGreeting(int) string
	// list goes on ...
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

// func printGreeti ng(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// very custom logic
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// very custom logic
	return "Hola!"
}
