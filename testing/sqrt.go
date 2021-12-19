// package main
package sqrt

import (
	"errors"
	"log"
)

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return float64(-1), errors.New("Invalid number! Cannot find sqrt of negative number.")
	}

	if num == float64(0) {
		return float64(0), nil
	}

	start := float64(0)
	end := num
	mid := float64(-1)
	for end-start >= 0.000001 {
		mid = (start + end) / 2
		log.Print("Mid: ", mid)

		// Square of mid number less than number then sqrt should be more than start.
		if mid*mid < num {
			start = mid
		}

		// Square of mid is greater than end then sqrt should be less then mid.
		if mid*mid >= num {
			end = mid
		}
	}
	return mid, nil
}

// func main() {
// 	sq, err := sqrt(14)
// 	if err != nil {
// 		log.Fatal("Error: ", err)
// 	}
// 	log.Print("Square root: ", sq)
// }
