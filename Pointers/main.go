package main

import "fmt"

func main() {
	var f float64
	f = 4.5

	address := &f
	value := *address

	fmt.Printf("Address of f: %v \nValue of f (using address): %v \n", address, value)

	// update value of f using address of f
	value = 6.5
	fmt.Printf("Value of modified f: %v \n", value)
}
