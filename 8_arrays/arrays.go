package main

import "fmt"

// numbered sequence of specific length
func main() {
	// int -> 0, string-> "", boolean -> false

	// zeroed values
	var nums [4]int
	// array length
	fmt.Println(len(nums))
	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)

	// false values
	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	// empty string
	var names [3]string
	names[0] = "golang"
	fmt.Println(names)

	// to declare it in single line
	numsDec := [3]int{1,2,3}
	fmt.Println(numsDec)

	// 2D arrays
	nums2d := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums2d)

	// fixed size, that is predictable
	// memory optimization
	// constant time access
}
