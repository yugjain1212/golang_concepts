package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	//for infinite loop
	// for {
	// 	fmt.Println("hello")
	// }
	//classic for loops
	for j := 1; j <= 20; j++ {
		if j%2 == 0 {
			continue
		} else if j > 14 {
			break
		}
		fmt.Println(j)
	}
	for i := range 5 {
		fmt.Println(i)
	}
}
