package main

import "fmt"

func main() {
	var x int
	x = 1

	var y int
	y = 2

	fmt.Printf("x=%v, type of x=%T", x, x)
	fmt.Println()
	fmt.Printf("y=%v, type of y=%T", y, y)
	fmt.Println()

	main := (x + y) / 2.0
	fmt.Printf("The avg of x and y is %v", main)
	fmt.Println()
}
