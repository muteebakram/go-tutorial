package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divMod(x int, y int) (int, int) {
	return x / y, x % y
}

func main() {
	x, y := 20, 10

	fmt.Println("x, y = ", x, y)
	div, mod := divMod(x, y)
	fmt.Println("add: ", add(x, y))
	fmt.Println("subtract: ", subtract(x, y))
	fmt.Println("multiply: ", multiply(x, y))
	fmt.Println("Div: ", div)
	fmt.Println("Mod: ", mod)
}
