package main

import (
	"log"
	"time"
)

type multipleString struct {
	s1 string
	s2 string
}
type myString struct {
	FirstName string
}

func (m *myString) printFirstName() string {
	return m.FirstName
}

type User struct {
	firstName string
	lastName  string
	age       int
	birthdate time.Time
}

func main() {
	var myStruct multipleString
	myStruct.s1 = "12"
	myStruct.s2 = "45"
	log.Println(saySomething(myStruct.s1))

	user := User{
		age:       15,
		firstName: "Hamza",
		lastName:  "Zeamari",
		birthdate: time.Now(),
	}

	log.Println(user)

	var myVar myString
	myVar.FirstName = "Hamza"

	myVar2 := myString{FirstName: "John"}
	log.Println("My var is set to : ", myVar.printFirstName())
	log.Println("My var2 is set to : ", myVar2.printFirstName())

}

func saySomething(s string) (string, string) {
	return s, "World"
}
