# Looping

Go supports only the `for` loop construct, but it supports mulitple flavors of the same:

## Normal for

```go
n := 10
for i := 0; i < n; i++ {
  fmt.Println("The current values is: ", i)
}
```

## Continue to skip an iteration
```go
n := 10
for i := 0; i < n; i++ {
  if i == 5 {
    continue
  }
  fmt.Println("The current values is: ", i)
}
```

## Break to break away from the loop
```go
n := 10
for i := 0; i < n; i++ {
  if i > 5 {
    break
  }
  fmt.Println("The current values is: ", i)
}
```

## For with only condition; init and increment/decrement as separate statements
```go
n := 10
i := 0
for i < n {
  fmt.Println("The current values is: ", i)
  i++
}
```

## Just the for; init, condition and increment/decrement as separate
```go
y := 0
for {
  if y == 5 {
    break
  }
  fmt.Println("The current values is: ", y)
  y++
}
```