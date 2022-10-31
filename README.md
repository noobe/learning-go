# Learning GO

01. Hello World
A basic program to print 'Hello, World!' in GO
- Create a file with extention `*.go`
- Declare a package name using syntax `package <name>` to group the functions that follows. The functions of the current package can be imported to other files using the `import "<name>"` synctax.
- Here we name the package as `main` and also fetch the `fmt` package to get the Println method from within to print to console

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

