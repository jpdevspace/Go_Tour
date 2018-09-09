package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is:", rand.Intn(10))
	fmt.Printf("The number %g is my favorite number\n", math.Sqrt(49))
}
