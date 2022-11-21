# 02. Variables and types.

## Go supports 2 ways to use variables:
- Declare variable and then assign a valid values
```go
var x int
x = 1
```
- Declare and and assign in a single line. Here Go makes type inference by checking the type of value being assigned to the variable.
```go
y := 2.0
```
- Setting multiple variables in the same line, but they all have to be of the same type.
```go
  x, y:= 3.0, 2.0
```
- Unused variables will result in compilation error.

- Arithmetic and logical operations can be performed only between values of same type
```go
x:=5
y:=6.0
z:= x+y // invalid operation: cannot compare x > y (mismatched types float64 and int)
```

## Go supports 10 int types 5 signed and 5 unsigned:
  int, uint
  int8, uint8
  int16, uint16
  int32, uint32
  int64, int64

## Go supports 2 types if floating points:
  float32, float64

## Printing to console:
  Println - Print unformatted normal string
  ```go
  fmlt.Println("This is the number: ", x)
  ```
  Printf - Print formatted text
  ```go
  fmlt.Printf("This is the number: %v", x)
  ```