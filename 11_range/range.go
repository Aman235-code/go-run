package main

import (
	"fmt"
	
)

// iteratig over data structures
func main() {
	// nums := []int{6, 7, 8}

	// for i := 0; i< len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println(i)
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	// m:= map[string]string{"fmname": "john", "lname": "doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// for k := range m {
	// 	fmt.Println(k) // key (default)
	// }

	// c- unicode code point rune
	// i -> starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}

}
