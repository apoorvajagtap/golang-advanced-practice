package main

import "fmt"

func main() {
	// Create an array of length 5 with values 45, 46, 67, 88, 99
	// add the values of on the index 2 and 4 and print the result
	arr := [5]int{45, 46, 67, 88, 99}
	sum := arr[2] + arr[4]

	fmt.Println("Sum of values on index 2 & 4: ", sum)

	// Create an array as above and print the slice a[i+3, j],
	// where i is starting index and j is last index
	fmt.Println("Printing the slice: ", arr[3:])
}
