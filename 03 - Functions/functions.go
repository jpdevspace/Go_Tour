package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// When two or more args share the same type
// we can omit the type on all except the last
func addMultiFive(a, b, c, d, e int) int {
	return a + b + c + d + e
}

func main() {
	fmt.Println(add(42, 11))
	fmt.Println(addMultiFive(7, 9, 1, 0, 10))
}
