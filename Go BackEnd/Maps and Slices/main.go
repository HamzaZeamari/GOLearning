package main

import (
	"fmt"
	"sort"
)

type User struct {
	firstname string
	lastname  string
}

func main() {
	
	var mySlice []string

	mySlice = append(mySlice, "First String")
	mySlice = append(mySlice, "Second String")
	mySlice = append(mySlice, "Third String")

	//Trier la slice/tableau
	sort.Strings(mySlice)

	myMap := make(map[string]string)
	myMapInt := make(map[string]int)
	myMapUser := make(map[string]User)

	myMap["Hamza"] = "Zeamari"
	myMapInt["One"] = 1
	myMapUser["Moi"] = User{firstname: "Hamza", lastname: "Zeamari"}

	fmt.Println(myMap["Hamza"])
	fmt.Println(myMapInt["One"])
	fmt.Println(myMapUser["Moi"])
	fmt.Println(mySlice)

}
