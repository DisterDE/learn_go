package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//alex := person{"Alex", "Anderson"}
	var alex person
	fmt.Println(alex)
}
