package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	urls := []string{
		"https://google.com/",
		"https://amazon.com/",
		"https://golang.org/",
		"https://stackoverflow.com/",
	}

	for _, url := range urls {
		go checkUrl(url)
	}
}

func checkUrl(url string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "is not available!")
		os.Exit(1)
	}

	fmt.Println(url, "is up!")
}
