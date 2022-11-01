package main

import "fmt"

func main() {
	msg := "This is a string."
	fmt.Println(msg)
	fmt.Println("Length: ", len(msg))
	fmt.Printf("Type is: %T \n", msg)
	fmt.Println("Char at 3rd position is: ", msg[3])
	fmt.Printf("Type of char at 3rd position: %T \n", msg[3])
	fmt.Println(msg + " And this is a concatenated string.")
	fmt.Println("Sliced from 3rd char till 10th char: ", msg[3:11])
	fmt.Println("Sliced from START till 10th char: ", msg[:11])
	fmt.Println("Sliced from 3rd till END char: ", msg[3:])
	rawString := `
		This is a 
		multiline
		raw string
	`
	fmt.Println(rawString)
}
