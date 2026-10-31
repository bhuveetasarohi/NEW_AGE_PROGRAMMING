package main

import "fmt"

func square(ch chan int) {
	fmt.Println("Square goroutine started")
	ch <- 5 * 5
	fmt.Println("Square goroutine completed")
}

func cube(ch chan int) {
	fmt.Println("Cube goroutine started")
	ch <- 3 * 3 * 3
	fmt.Println("Cube goroutine completed")
}

func fibonacci(ch chan int) {
	fmt.Println("Fibonacci goroutine started")

	a, b := 0, 1
	for i := 0; i < 6; i++ {
		a, b = b, a+b
	}

	ch <- a
	fmt.Println("Fibonacci goroutine completed")
}

func main() {
	ch := make(chan int)

	go square(ch)
	go cube(ch)
	go fibonacci(ch)

	for i := 0; i < 3; i++ {
		fmt.Println("Result:", <-ch)
	}

	fmt.Println("All tasks completed")
}
