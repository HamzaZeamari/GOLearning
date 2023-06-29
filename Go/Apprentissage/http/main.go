package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println("Error")
		os.Exit(1)
	}
}
