package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)
	for _, link := range links {
		time.Sleep(1 * time.Second)
		go checkLink(link, channel)
	}

	for l := range channel {
		go checkLink(l, channel)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
		return
	}

	fmt.Println(link, "is up!")
	channel <- link
}
