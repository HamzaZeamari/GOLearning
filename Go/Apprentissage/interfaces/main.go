package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	english := englishBot{}
	spanish := spanishBot{}

	printGreeting(english)
	printGreeting(spanish)
}

func (spanishBot) getGreeting() string {
	return "Holà !"
}
func (englishBot) getGreeting() string {
	return "Hi there !"
}
func printGreeting(B Bot) {
	fmt.Println(B.getGreeting())
}
