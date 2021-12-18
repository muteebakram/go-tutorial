package main

import "fmt"

func main() {
	x := 10
	if x > 10 {
		fmt.Println("Value is greater than 10")
	} else if x == 10 {
		fmt.Println("Value is equal 10")
	} else {
		fmt.Println("Value is less than 10")
	}

	if frac := 10 % 2; frac == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd Number")
	}
}
