package main

import (
	"fmt"
	"math"
)

// interface
type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) perimeter() float64 {
	return 4 * s.length
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShape(s shape) {
	fmt.Println("Area of the shape:",s.area())
	fmt.Println("Perimeter of the shape:",s.perimeter())
}

func main(){
	shapes:= []shape{
		square{length: 15.4},
		circle{radius: 10.3},
		square{length: 12},
		circle{radius: 13},
	}

	for _,v := range shapes{
		printShape(v)
		fmt.Println("---")
	}
}
