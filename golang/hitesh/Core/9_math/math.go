package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Floor and Ceil
	fmt.Println("Floor of 3.8:", math.Floor(3.8)) // Output: 3
	fmt.Println("Ceil of 3.8:", math.Ceil(3.8))   // Output: 4

	// Log and Log10
	fmt.Println("Natural logarithm of 10:", math.Log(10))     // Output: 2.302585092994046
	fmt.Println("Base-10 logarithm of 100:", math.Log10(100)) // Output: 2

	// Round
	fmt.Println("Round of 3.5:", math.Round(3.5)) // Output: 4
	fmt.Println("Round of 3.4:", math.Round(3.4)) // Output: 3

	// Absolute Value
	fmt.Println("Absolute value of -5:", math.Abs(-5)) // Output: 5

	// Power
	fmt.Println("2^3:", math.Pow(2, 3)) // Output: 8

	// Min and Max
	fmt.Println("Minimum of 5 and 10:", math.Min(5, 10)) // Output: 5
	fmt.Println("Maximum of 5 and 10:", math.Max(5, 10)) // Output: 10

	// Generate randomNumber

	rand.Seed(time.Now().UnixNano())

	max := 77
	min := 23

	randomNumber := rand.Intn(max-min+1) + min

	fmt.Println(randomNumber)
}
