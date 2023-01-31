package main

import "fmt"

func main() {
	// Convert the float value 67.88 into int
	fVal := 67.88
	iVal := int(fVal)
	fmt.Println("int converted value of '67.88': ", iVal)

	// Create 3 variables num1 (int), num2(float64), sum(float64), Add num1 and num2
	num1, num2, sum := 5, 6.1, 0.0
	sum = float64(num1) + num2
	fmt.Println("Sum of num1 & num2: ", sum)
}
