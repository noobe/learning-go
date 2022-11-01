package main

import "fmt"

func main() {
	n := 10

	fmt.Println("Looping from 0 to 9")
	for i := 0; i < n; i++ {
		fmt.Println("The current values is: ", i)
	}
	fmt.Println("------------------")

	fmt.Println("Looping from 0 to 9 but breaking off at 5")
	for i := 0; i < n; i++ {
		if i > 5 {
			break
		}
		fmt.Println("The current values is: ", i)
	}
	fmt.Println("------------------")

	fmt.Println("Looping from 0 to 9 but skipping 5")
	for i := 0; i < n; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("The current values is: ", i)
	}
	fmt.Println("------------------")

	fmt.Println("Other formats of for loop")
	i := 0
	for i < n {
		fmt.Println("The current values is: ", i)
		i++
	}
	fmt.Println("------------------")

	fmt.Println("Other formats of for loop")
	y := 0
	for {
		if y == 5 {
			break
		}
		fmt.Println("The current values is: ", y)
		y++
	}
	fmt.Println("------------------")
}
