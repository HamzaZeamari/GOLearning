package main

import "fmt"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "First String")
	mySlice = append(mySlice, "Second String")
	mySlice = append(mySlice, "Third String")

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}
	for _, stringValue := range mySlice {
		fmt.Println(stringValue)
	}
}
