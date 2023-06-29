package main

import "fmt"

func main() {
	if 1 == 2 {
		fmt.Println(true)
	} else if 1 == 3 {
		fmt.Println(false)
	} else {
		fmt.Println("OK")
	}
	cat := "Yellow"

	switch cat {
	case "Red":
		fmt.Println("red")
	case "Yellow":
		fmt.Println("Yellow")
	default:
		fmt.Println("Another color")
	}
}
