package main

import "fmt"

func main() {
	// Create an map for storing 5 student marks (type : map[string]int{}
	marks := map[string]int{
		"A": 50,
		"B": 20,
		"C": 30,
		"D": 34,
		"E": 34,
	}

	// Check if any student have 34 marks increment it to 1 and
	// then print the map(Iterate over the maps)
	for k, v := range marks {
		if v == 34 {
			marks[k] += 1
		}
	}
	fmt.Println(marks)
}
