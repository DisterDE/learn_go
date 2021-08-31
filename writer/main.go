package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(logWriter{}, resp.Body)
}

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes were written:", len(bs))
	return len(bs), nil
}
