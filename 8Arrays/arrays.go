package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(len(numbers), numbers[0], numbers[1], numbers[2])

	var numbers2 [3]int = [3]int{1, 2, 3}
	fmt.Println(len(numbers2), numbers2[0], numbers2[1], numbers2[2])
	var char [3]string = [3]string{"a", "b", "c"}
	fmt.Println(len(char), char[0], char[1], char[2])
	//var numbers3 []int = [3]int{1, 2, 3} //error: cannot use array (type [3]int) as type []int in assignment
	nums := [3][3]int{{3, 5, 4}, {1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)
	for i := range nums {
		for j := range nums[i] {
			fmt.Printf("%d ", nums[i][j])
		}
		fmt.Println()
	}

}
