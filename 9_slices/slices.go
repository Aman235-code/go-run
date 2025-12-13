package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums==nil)
	fmt.Println(len(nums))

	var nums2 = make([]int, 2, 5)
	fmt.Println(nums2)
	fmt.Println(nums2==nil) // false
	fmt.Println(cap(nums2)) // capacity is 5 -> max nos of elements can fit
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	fmt.Println(len(nums2))

	nums3 := []int{}
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 2)
	fmt.Println(nums3)
	fmt.Println(cap(nums3))
	fmt.Println(len(nums3))

	var nums4 = make([]int, 2, 5)
	nums4[0] = 3
	nums4[1] = 5
	fmt.Println(nums4)
	fmt.Println(cap(nums4))
	fmt.Println(len(nums4))

	// copy function
	var nums5 = make([]int, 0, 5)
	nums5 = append(nums5, 2)
	var nums6 = make([]int, len(nums5))

	fmt.Println(nums5)
	fmt.Println(nums6)

	copy(nums6, nums5)
	fmt.Println(nums5)
	fmt.Println(nums6)

	// slice operator
	var nums7 = []int{1, 2, 3, 4, 5}
	fmt.Println(nums7[0:2])
	fmt.Println(nums7[:2])
	fmt.Println(nums7[1:])

	// slice package
	var nums8 = []int{1, 2}
	var nums9 = []int{1, 2}
	fmt.Println(slices.Equal(nums8, nums9)) // true

	var nums10 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums10)

}
