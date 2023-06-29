package main

import "log"

func main() {
	var myString string
	myString = "hamza"

	log.Println("MyString is set to ", myString)
	changeUserPointer(&myString)
	log.Println("MyString is set to ", myString, " after modification ")

}

func changeUserPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
