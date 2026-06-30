package main

import "fmt"

func changenum(num *int) {
	*num = 5
	fmt.Println("the numchange", num)

}
func main() {
	num := 1
	changenum(&num)
	fmt.Println("the num", num)
}
