package main

import "fmt"

// Create an interface to calculate the area of Square and Rectangle
// Implement the Shape interface using type Square and rectangle
type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (s Square) Area() float64 {
	return (s.Side * s.Side)
}

func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

func main() {
	rA := Rectangle{Height: 4, Width: 6}
	sA := Square{Side: 10}

	fmt.Printf("Area of Rectangle: %v\nArea of Square: %v\n", rA.Area(), sA.Area())
}
