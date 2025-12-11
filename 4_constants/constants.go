package main

import "fmt"

const age1 = 34

var name1 = "aman"

// name:= "aman" // not allowed

func main() {
	const name string = "golang"
	// name = "js" // error
	const age = 30

	fmt.Println(age1)
	fmt.Println(name1)

	const (
		port = 5000
		host = "localhost"
	)

	// port = 5500 // error 

	fmt.Println(port, host)
}