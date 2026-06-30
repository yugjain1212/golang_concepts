package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Department string
	Salary     float64
}

func main() {
	e := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		Department: "Engineering",
		Salary:     60000.0,
	}

	fmt.Println(e.Name)       // prints "John Doe"
	fmt.Println(e.Age)        // prints 30
	fmt.Println(e.Department) // prints "Engineering"
	fmt.Println(e.Salary)     // prints 60000
}
