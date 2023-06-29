package main

import "fmt"

func main() {

	whatToSay, whatToDo := saySomething()
	var i int = 15

	fmt.Printf("i is set to %d \n", i)
	fmt.Printf("Here's what to say: %s and what to do: %s \n", whatToSay, whatToDo)
}

func saySomething() (string, string) {
	return "Something", "Anything"
}
