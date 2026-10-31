
package main

import "fmt"

// Function that modifies the original variable using pointer
func changeValue(ptr *int) {
    *ptr = 100
}

// Structure
type Student struct {
    Name string
    Age  int
}

func main() {

    // Part 1: Pointer declaration, address and value
    num := 10

    ptr := &num

    fmt.Println("Part 1: Basic Pointer")
    fmt.Println("Original Value:", num)
    fmt.Println("Address of num:", ptr)
    fmt.Println("Value using pointer:", *ptr)

    // Part 2: Modifying original variable
    fmt.Println("\nPart 2: Pass by Reference")
    fmt.Println("Before modification:", num)

    changeValue(ptr)

    fmt.Println("After modification:", num)

    // Part 3: Using new() with structure
    fmt.Println("\nPart 3: Structure Pointer")

    student := new(Student)

    student.Name = "Rahul"
    student.Age = 20

    fmt.Println("Student Name:", student.Name)
    fmt.Println("Student Age:", student.Age)

    // Modifying structure fields through pointer
    student.Age = 21

    fmt.Println("Updated Student Age:", student.Age)
}
