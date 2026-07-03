package main

import "fmt"

func printslice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	items []T
}

func main() {
	// var items []int = []int{1, 2, 3, 4, 5}
	itams := []int{1, 2, 3, 4, 5}
	printslice(itams)
	items := []string{"one", "two", "three", "four", "five"}
	printslice(items)
	mystack := stack[int]{items: []int{1, 2, 3, 4, 5}}
	fmt.Println(mystack)
}
