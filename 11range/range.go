package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// 	sum = sum + num
	// }
	// fmt.Println(sum)
	// m := map[string]int{"a": 1, "b": 2, "c": 3}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	//unicode point
	// starting byte of run
	//255 is the max value of a byte in golang
	for i, c := range "hello world" {
		fmt.Println(i, c)
	}
}
