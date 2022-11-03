package main

import "fmt"

func main() {
	strSlice := []string{"str1", "str2", "str3"}
	numSlice := []int{1, 2, 3, 4}

	fmt.Printf("String slice: %v, type: %T\n", strSlice, strSlice)
	fmt.Printf("Integer slice: %v, type: %T\n", numSlice, numSlice)

	fmt.Println("Length of String slice: ", len(strSlice))
	fmt.Println("Length of Integer slice: ", len(numSlice))

	fmt.Println("The 2nd element in the slice is :", strSlice[1])
	fmt.Println("The 2nd element in the slice is :", numSlice[1])

	fmt.Println("A smaller slice be sliced from position 1-3 of parent slice", strSlice[1:3])
	fmt.Println("A smaller slice be sliced from position 1-3 of parent slice", numSlice[1:3])

	fmt.Println("Iterating the slice using a for loop:")
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}

	fmt.Println("Using range and only index:")
	for i := range numSlice {
		fmt.Println(i, numSlice[i])
	}

	fmt.Println("Using range and index & value:")
	for i, name := range strSlice {
		fmt.Println(i, name)
	}

	fmt.Println("Using range and only value:")
	for _, name := range numSlice {
		fmt.Println(name)
	}

	fmt.Println("Appending to slice:")
	newStrSlice := append(strSlice, "Test")
	fmt.Println(strSlice, newStrSlice)
}
