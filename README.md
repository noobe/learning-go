# Learning GO
Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

## Writing and executing Go code
A basic program to print 'Hello, World!' in GO.
- Create a file with extention `*.go`.
- Every Go code lives inside a package.
- All the files in the same folder in the disc should be under the same package.
- Declare a package name using syntax `package <name>` to group the functions that follows. The functions of the current package can be imported to other files using the `import "<name>"` synctax.
- Here we name the package as `main`. `main` is a special package the compiles the code into an executable rather than treating it as a lib.
- We also fetch the `fmt`(formatted) package to get the `Println` method from within to print to console. Go keeps it simple we need to import other packages for additional functionality.
- To use functions from other packages, we need to refer it as `<package-name>.<function>()`.
- Every function in an exported package should start with an Upper casing. Eg: fmt.Println()
- Here we write our code in the `main()` function. This is the function that gets triggered when we run a go executable and hence it should be a unique function in the project package.
- No semicolons in go.
- To build: we call `go build <file-name>.go` from its folder and it creates an executable in the same folder named `<file-name>`.
- To Run: we can call the output `./<file-name>` directly as its an executable.
- To Build and Run it in one line - we call `go run <file-name>.go` from its folder.
Eg:
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

## go help
Usage of go tools:
  go <command> [arguments]

The commands are:
  bug         start a bug report
  build       compile packages and dependencies
  clean       remove object files and cached files
  doc         show documentation for package or symbol
  env         print Go environment information
  fix         update packages to use new APIs
  fmt         gofmt (reformat) package sources
  generate    generate Go files by processing source
  get         add dependencies to current module and install them
  install     compile and install packages and dependencies
  list        list packages or modules
  mod         module maintenance
  work        workspace maintenance
  run         compile and run Go program
  test        test packages
  tool        run specified go tool
  version     print Go version
  vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:
  buildconstraint build constraints
  buildmode       build modes
  c               calling between Go and C
  cache           build and test caching
  environment     environment variables
  filetype        file types
  go.mod          the go.mod file
  gopath          GOPATH environment variable
  gopath-get      legacy GOPATH go get
  goproxy         module proxy protocol
  importpath      import path syntax
  modules         modules, module versions, and more
  module-get      module-aware go get
  module-auth     module authentication using go.sum
  packages        package lists and patterns
  private         configuration for downloading non-public code
  testflag        testing flags
  testfunc        testing functions
  vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.

## Topics covered:
01. [Hello World](./01%20-%20Hello%20World)
02. [Variables and assignments](./02%20-%20Variables%20and%20assignments)
03. [Conditionals](./03%20-%20Conditionals)
04. [Looping](./04%20-%20Looping)
05. [Code challenge combining Looping and branching](./05%20-%20Challlenge%20-%20Fizz%20buzz)
06. [Strings](./06%20-%20Strings)
07. [Code challenge](./07%20-%20Challenge%20-%20Even-ended-number)
08. [Arrays](./08%20-%20Arrays)
09. [Slices](./09%20-%20Slices)
10. [Maps](./10%20-%20Maps)
11. [Functions](./11%20-%20Functions)
12. [Structs](./12%20-%20Structs)
