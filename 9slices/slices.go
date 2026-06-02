package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic arrays
// most use construct go //usefull method
func main() {
	var nums = make([]int, 2, 5)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))
	fmt.Println(nums)
	var nums2 = make([]int, 2, 5)
	var nums3 = make([]int, len(nums2))
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	copy(nums3, nums2)
	fmt.Println(nums2, nums3)

	//slice operator
	var nums4 = []int{1, 2, 3, 4, 5}
	var nums5 = []int{6, 7, 8, 9, 10}

	fmt.Println(nums4[0:2])
	fmt.Println(nums4[:3])
	fmt.Println(nums4[2:])
	fmt.Println(nums4[:])
	fmt.Println(slices.Equal(nums4, nums5))
	var nums6 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums6)
}
