package main

import "fmt"

func main() {
	x := 5.0

	if x > 5 {
		fmt.Println("X is larger than 5")
	}

	if x > 100 {
		fmt.Println("X is larger than 100")
	} else {
		fmt.Println("X is smaller than 100")
	}

	if x > 5 && x < 10 {
		fmt.Println("X is between 5 and 10")
	} else {
		fmt.Println("X is not between 5 and 10")
	}

	if x < 5 || x > 10 {
		fmt.Println("X is outside the range of 5 to 10")
	} else {
		fmt.Println("X is inside the range of 5 to 10")
	}

	if y := 6.0; x > y {
		fmt.Println("X is bigger than newly set value y")
	} else {
		fmt.Println("X is smaller than newly set value y")
	}

	n := 3

	switch n {
	case 1:
		fmt.Println("n is One")
	case 2:
		fmt.Println("n is Two")
	case 3:
		fmt.Println("n is Three")
	case 4:
		fmt.Println("n is Four")
	case 5:
		fmt.Println("n is Five")
	default:
		fmt.Println("n is something else")
	}
}
