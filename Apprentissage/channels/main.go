package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urlSites := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.apple.com",
		"http://www.facebook.com",
		"http://twitter.com",
	}

	channel := make(chan string)

	for _, site := range urlSites {
		go checkLink(site, channel)
	}

	for l := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(l)
	}

}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " est OUT")
		channel <- link
	} else {
		fmt.Println(link, " est IN")
		channel <- link

	}
}
