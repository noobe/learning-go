# Conditionals
Go supports these constructs for branching in logic:
1. if
```go
if x > 5 {
  fmt.Println("X is larger than 5")
}
```
2. if else
```go
if x > 100 {
  fmt.Println("X is larger than 100")
} else {
  fmt.Println("X is smaller than 100")
}
```
3. else if
```go
if x > 100 {
  fmt.Println("X is larger than 100")
} else if x > 50 {
  fmt.Println("X is larger than 50")
} else {
  fmt.Println("X is smaller than 50")
}
```
4. switch
```go
switch n {
  case 1:
    fmt.Println("n is One")
  case 2:
    fmt.Println("n is Two")
  case 3:
    fmt.Println("n is Three")
  case 4:
    fmt.Println("n is Four")
  case 5:
    fmt.Println("n is Five")
  default:
    fmt.Println("n is something else")
}
```