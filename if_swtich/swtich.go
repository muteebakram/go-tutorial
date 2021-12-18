package main

func main() {
	x := 10
	switch {
	case x > 100:
		println("Greater")
	case x < 100:
		println("Lesser")
	case x == 100:
		println("Equal")
	default:
		println("Unkown")

	}
}
