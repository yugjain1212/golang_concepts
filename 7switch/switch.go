package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	fmt.Print("enter a number: ")
	fmt.Scanln(&i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its a weekend")
	default:
		fmt.Println("its a weekday")
	}

	//type switch
	whoAmi := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its an int")
		case string:
			fmt.Println("its a string")
		default:
			fmt.Println("i dont know", t)
		}
	}
	whoAmi(1)
	whoAmi("yug")
	whoAmi(true)

}
