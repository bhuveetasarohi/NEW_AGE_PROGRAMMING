package main

import (
	"fmt"
	"math"
)

// Error 1: Division by zero
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Error 2: Negative square root
func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative number not allowed")
	}
	return math.Sqrt(n), nil
}

// Error 3: Empty input
func checkInput(s string) error {
	if s == "" {
		return fmt.Errorf("input is empty")
	}
	return nil
}

func main() {
	result, err := divide(10, 2)
	fmt.Println("Division:", result, err)

	result, err = squareRoot(-9)
	fmt.Println("Square Root:", result, err)

	err = checkInput("")
	fmt.Println("Input:", err)
}
