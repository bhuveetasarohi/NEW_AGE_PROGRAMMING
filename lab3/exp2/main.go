package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) ReadData() {
	fmt.Print("Enter Name: ")
	fmt.Scan(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scan(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&p.Salary)
}

func (p Person) DisplayData() {
	fmt.Println("\nPerson Details")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Job:", p.Job)
	fmt.Printf("Salary: %.2f\n", p.Salary)
}

func main() {

	var person1 Person

	fmt.Println("Enter details of First Person:")
	person1.ReadData()
	person1.DisplayData()

	var person2 Person

	fmt.Println("\nEnter details of Second Person:")
	person2.ReadData()
	person2.DisplayData()
}
