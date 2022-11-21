# 01. Hello World.
1. Go code is saved in files with extension `*.go`

2. Go file consists of:
  - Package declarations
  - Import packages
  - Functions
  - Statements and expressions

3. Code in file:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
- In Go every program is part of a package, which we define using the `package` keyword. In the current example, we name the package as `main`
- We are importing the `fmt` package next using the `import` keyword followed by the package name. Package `fmt` implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
- The `func` keyword is used to declare a function.
- The `Println()` method available from the `fmt` package allows us to Print the given string to a newline in the console.
- Statements can be ended using the Enterkey or the `;`
- Blocks of code are grouped using the `{` & `}` but a line can not be started using a `{`.
```go
package main
import ("fmt")

func main()
{
  fmt.Println("Hello World!")
}
```
Give the following error message
```go
./prog.go:4:6: missing function body
./prog.go:5:1: syntax error: unexpected semicolon or newline before {
```
- Go code is compiled and run using the syntax `go run <file-name>`