package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com/",
		"https://amazon.com/",
		"https://golang.org/",
		"https://stackoverflow.com/",
		"https://facebook.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkUrl(link, c)
		}(l)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
