package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
}

type Square struct {
	x      float64
	y      float64
	length float64
}

type Circle struct {
	radius float64
}

func NewSquare(x float64, y float64, length float64) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Invalid length: %f", length)
	}

	return &Square{x: x, y: y, length: length}, nil
}

func NewCirlce(radius float64) (*Circle, error) {
	if radius <= 0 {
		return nil, fmt.Errorf("Invalid radius: %f", radius)
	}

	return &Circle{radius: radius}, nil
}

func (s *Square) Area() float64 {
	return s.length * s.length
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func areaSum(sh []Shape) float64 {
	sum := float64(0)
	for _, s := range sh {
		fmt.Println("Area of ", reflect.TypeOf(s), ": ", s.Area())
		sum += s.Area()
	}
	return sum
}

func main() {
	fmt.Println("Struct example:")
	sq, err := NewSquare(0, 0, 2)
	if err != nil {
		fmt.Printf("Error: Failed to initailize square. Reason: %s\n", err)
		return
	}
	fmt.Printf("Square defined: %v\n", sq)

	c, err := NewCirlce(2)
	if err != nil {
		fmt.Printf("Error: Failed to initailize square. Reason: %s\n", err)
		return
	}
	fmt.Println("Circle defined: ", c)

	shapes := []Shape{sq, c}
	fmt.Println("Sum of areas: ", areaSum(shapes))
}
