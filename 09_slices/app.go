package main

import "fmt"

// slice -> dynamic
// most used construct in go
// + useful methods

func main() {
	// uninitialized slice is nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// capacity -> maximum numbers of elements can fit
	// nums :=[]int{}
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy function
	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	// slice operator

	var nums = []int{1,2,3}
	fmt.Println(nums[0:2])
}
