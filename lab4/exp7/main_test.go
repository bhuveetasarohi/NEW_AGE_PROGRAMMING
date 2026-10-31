package main

import "testing"

// Test 1: Division
func TestDivide(t *testing.T) {
	result, err := divide(10, 2)

	if err != nil || result != 5 {
		t.Error("Division test failed")
	}

	_, err = divide(10, 0)

	if err == nil {
		t.Error("Division by zero test failed")
	}
}

// Test 2: Square root
func TestSquareRoot(t *testing.T) {
	result, err := squareRoot(25)

	if err != nil || result != 5 {
		t.Error("Square root test failed")
	}

	_, err = squareRoot(-9)

	if err == nil {
		t.Error("Negative square root test failed")
	}
}

// Test 3: Empty input
func TestCheckInput(t *testing.T) {
	err := checkInput("Hello")

	if err != nil {
		t.Error("Valid input test failed")
	}

	err = checkInput("")

	if err == nil {
		t.Error("Empty input test failed")
	}
}
