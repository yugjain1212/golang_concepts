package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func language() (string, string, string) {
	return "python", "java", "go"

}

//	func processit(fn func(a int)int){
//		fn(10)
//	}
func processit() func(o int) int {
	return func(o int) int {
		return 4
	}
}
func main() {
	var p, q int
	p = 10
	q = 20
	fmt.Println(add(p, q))
	fmt.Println(sub(p, q))
	fmt.Println(language())
	fn := processit()
	fmt.Println(fn(10))

}
