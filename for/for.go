package main

func main() {
	for i := 0; i < 10; i++ {
		println("i = ", i)
	}

	println("\nWhile loop using for")
	i := 0
	for i < 10 {
		println("i = ", i)
		i++
	}

	println("\nWhile of true")
	for {
		if i == 20 {
			break
		}
		println("i = ", i)
		i++
	}

	println("\nRange in go")
	fruits := []string{"Apple", "Banana", "Orange"}
	for index, fruit := range fruits {
		println("Index: ", index, "Fruits: ", index, fruit)
	}

	println("\nMax value in go")
	x := []int{1, 2, 3, 4, 5, 6, 7}
	max := x[0]
	for _, number := range x[1:] {
		if number > max {
			max = number
		}
	}
	println("Max: ", max)

}
