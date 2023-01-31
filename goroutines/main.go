package main

import (
	"fmt"
	"sync"
)

// waitground
var wg sync.WaitGroup

// Create a struct of employee with data name and age
type Employee struct {
	Name string
	Age  int
	Data string
}

func sendEmpData(ch chan Employee, age int) {
	defer wg.Done()
	ch <- Employee{
		Name: "Apoorva",
		Age:  age,
	}
}

func main() {
	// Create a goroutine which Send the Employees data to the channel with increementing age.
	// In main function create a buffered channel and then call the 10 goroutines (use WaitGroups)

	ch := make(chan Employee, 10)
	ag := 24

	// Adding wg.Add() & wg.Wait() inside the for loop as we need it in sequence.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sendEmpData(ch, ag)
		ag++
		wg.Wait()
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	close(ch)
}
