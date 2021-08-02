package main

import "fmt"

func main() {
	numbers := newSlice()

	for _, n := range numbers {
		fmt.Print(n, " is ")
		if n%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}

func newSlice() []int {
	var result []int
	for i := 0; i <= 10; i++ {
		result = append(result, i)
	}
	return result
}
