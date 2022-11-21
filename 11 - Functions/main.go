package main

import (
	"fmt"
)

func myFunction() {
	fmt.Println("My Function")
}

func myFunctionWithParam(name string) {
	fmt.Println("Hello ", name)
}

func myFunctionWithParams(name string, age int) {
	fmt.Println("Hello ", name, ". You are ", age, " years old")
}

func myFunctionWithReturn() string {
	return "Hello!"
}

func main() {
	myFunction()
	myFunctionWithParam("Alex")
	myFunctionWithParams("Tom", 12)
	fmt.Println(myFunctionWithReturn())
}
