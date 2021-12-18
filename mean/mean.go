package main

import "fmt"

func main() {
	fmt.Printf("Calculating mean of two numbers\n")

	var x float32
	var y float32

	fmt.Printf("Default x=%v type=%T\n", x, x)
	fmt.Printf("Default y=%v type=%T\n", y, y)

	x = 1
	y = 2

	fmt.Printf("x=%v type=%T\n", x, x)
	fmt.Printf("y=%v type=%T\n", y, y)

	var mean = (x + y) / 2
	fmt.Printf("Mean=%v type=%T\n", mean, mean)
}
