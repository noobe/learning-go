package main

import "fmt"

func main() {
	count := 0
	for i := 1000; i <= 9999; i++ {
		for j := 1000; j <= 9999; j++ {
			n := i * j
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
				// fmt.Printf("# %d. %d\n", count, n)
			}
		}
	}
	fmt.Printf("There are %d even ended numbers in the possible results when we multiply 2 four digit numbers\n", count)
}
