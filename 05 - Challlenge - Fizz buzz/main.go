package main

import "fmt"

func main() {
	max := 20

	for i := 1; i <= max; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
			continue
		}
		if i%3 != 0 && i%5 == 0 {
			fmt.Println("buzz")
			continue
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
			continue
		}
		fmt.Println(i)
	}
}
