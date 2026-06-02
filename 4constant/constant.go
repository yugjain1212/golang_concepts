package main

import "fmt"

func main() {
	const pi float64 = 3.14
	const pi2 = 3.14
	const pi3 = 3.14
	fmt.Println(pi)
	fmt.Println(pi2)
	fmt.Println(pi3)
	const (
		pi4 = 3.145
		pi5 = 3.14
	)
	fmt.Println(pi4)
	fmt.Println(pi5)
}
