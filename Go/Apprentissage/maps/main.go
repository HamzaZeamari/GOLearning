package main

import "fmt"

func main() {
	//var color map[string]string
	//colorss := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#002BFF",
		"green": "#0CFF00",
	}

	// ADD
	colors["white"] = "#FFFFFF"
	// DELETE
	delete(colors, "white")

	print(colors)
}

func print(m map[string]string) {
	for i, color := range m {
		fmt.Println(i+" ", color)
	}
}
