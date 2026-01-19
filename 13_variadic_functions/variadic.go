package main

import "fmt"

// n number of integers
func sum(nums ...int) int {
	total:=0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {
	// fmt.Println(1, 2, 3, 4, 5, 88, "hello") // n number of arguments it can take

	nums := []int{3, 4, 5, 6}

	result:= sum(nums...)

	fmt.Println(result)
}