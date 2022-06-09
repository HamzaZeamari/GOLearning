package main

import (
	"fmt"
	"net/http"
)

func main() {
	urlSites := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.apple.com",
		"http://www.facebook.com",
		"http://twitter.com",
	}

	for _, site := range urlSites {
		checkLink(site)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " est OUT")
	} else {
		fmt.Println(link, " est IN")
	}
}
