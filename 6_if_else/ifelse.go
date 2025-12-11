package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("person is an adult")
	} else {
		fmt.Println("person is not an adult")
	}

	age1:= 16

	if age1 >= 18 {
		fmt.Println("person is an adult")
	} else if age1 >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions {
		fmt.Println("Yes")
	}

	if role == "admin" && hasPermissions {
		fmt.Println("Yes")
	}

	// we can declare a variable inside the if construct
	if age:= 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// goe does not have ternary operator have to use normal if else
}