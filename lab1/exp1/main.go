package main

import "fmt"

func main() {

	var a, b int
	var x, y float64

	// Integer input
	for {
		fmt.Print("Enter two integers: ")
		n, err := fmt.Scan(&a, &b)

		if err == nil && n == 2 && a >= 0 && b >= 0 && a <= 1000 && b <= 1000 {
			break
		}

		fmt.Println("Invalid input! Try again.")
	}

	// Integer operations
	fmt.Println("Integer Addition:", a+b)
	fmt.Println("Integer Subtraction:", a-b)
	fmt.Println("Integer Multiplication:", a*b)

	// Float input
	for {
		fmt.Print("\nEnter two decimal numbers: ")
		n, err := fmt.Scan(&x, &y)

		if err == nil && n == 2 && x >= 0 && y >= 0 && x <= 1000 && y <= 1000 {
			break
		}

		fmt.Println("Invalid input! Try again.")
	}

	// Float operations
	fmt.Println("Float Addition:", x+y)
	fmt.Println("Float Subtraction:", x-y)
	fmt.Println("Float Multiplication:", x*y)
}
