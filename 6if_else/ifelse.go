package main

import "fmt"

func main() {
	age := 18
	if age >= 18 && age <= 65 {
		fmt.Println("you can vote")
	} else if age == 17 {
		fmt.Println("you can learn to drive")
	} else {
		fmt.Println("you cannot vote")
	}
}
