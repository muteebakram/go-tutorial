package main

func main() {
	for number := 1; number <= 100; number++ {
		if number%3 == 0 && number%5 == 0 {
			println("fizzbuzz")
		} else if number%3 == 0 && number%5 != 0 {
			println("fizz")
		} else if number%3 != 0 && number%5 == 0 {
			println("buzz")
		} else {
			println(number)
		}
	}
}
