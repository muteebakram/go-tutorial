package main

import (
	"fmt"
	"math"
)

func squareRoot(num float64) (error, float64) {
	if num < 0.0 {
		return fmt.Errorf("Negative number."), 0.0
	}

	return nil, math.Sqrt(num)
}

func main() {
	err, val := squareRoot(-2.2)
	if err != nil {
		fmt.Printf("Failed to get value. Reason: %s", err)
	} else {
		fmt.Printf("Square Root is %f", val)
	}

}
