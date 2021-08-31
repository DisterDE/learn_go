package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name is not found in args")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println(err)
	}
}
