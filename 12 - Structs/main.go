package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary float32
}

func main() {
	var p1 Person
	var p2 Person

	p1.name = "Alice"
	p2.name = "Bob"

	p1.age = 14
	p2.age = 15

	p1.job = "Engineer"
	p2.job = "Tester"

	p1.salary = 10.1
	p2.salary = 9.9

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.name)
	fmt.Println(p2.name)
}
