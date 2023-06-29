package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	lastName  string
	firstName string
	contactInfo
}

func main() {

	slice := []string{"Hi", "There", "Hamza"}
	updateSlice(slice)
	fmt.Println(slice)

	// Méthodes pour personnes sans contactInfo
	// Méthode 1
	//pers := Person{"Hamza", "Zeamari"}
	// Méthode 2
	//pers := Person{firstName: "Hamza", lastName: "Zeamari"}
	// Méthode 3
	//var pers Person
	//pers.firstName = "Hamza"
	//pers.lastName = "Zeamari"
	/*
		pers := Person{
			firstName: "Hamza",
			lastName:  "Zeamari",
			contactInfo: contactInfo{
				email:   "hamza@gmail.com",
				zipCode: 67100,
			},
		}
	*/
	//pers.updateName("Karim")
	//pers.print()
}

func updateSlice(s []string) {
	s[0] = "Yop"
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateName(newFirstname string) {
	(*p).firstName = newFirstname
}
