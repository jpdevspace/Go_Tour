package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Notice how Intn imported from rand stards with capital 'I'
// Packages only export capitalized methods.

func main() {
	fmt.Println("My favorite number is:", rand.Intn(10))
	fmt.Printf("The number %g is my favorite number\n", math.Sqrt(49))
}
