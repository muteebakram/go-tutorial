package main

import "fmt"

type Square struct {
	x      float64
	y      float64
	length float64
}

func NewSquare(x float64, y float64, length float64) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Invalid length: %f", length)
	}

	return &Square{x: x, y: y, length: length}, nil
}

func Move(s *Square, x float64, y float64) *Square {
	if s == nil {
		fmt.Printf("Error: Invalid square cannot move. Sq: %v\n", s)
		return nil
	}
	s.x = x
	s.y = y

	return s
}

func Area(s *Square) float64 {
	return s.length * s.length
}

func main() {
	fmt.Println("Struct example:")
	sq, err := NewSquare(0, 0, 2)
	if err != nil {
		fmt.Printf("Error: Failed to initailize square. Reason: %s\n", err)
		return
	}
	fmt.Printf("Square defined: %v\n", sq)

	Move(sq, 1, 1)
	fmt.Printf("Square moved: %v, %v\n", sq.x, sq.y)
	fmt.Printf("Area of square: %f", Area(sq))
}
