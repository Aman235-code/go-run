package main

import (
	"fmt"
	"maps"
)

// maps => hash, objects, dict
func main() {
	// creating map
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"], m["area"])

	// imp: if key does not exist in the map then it returns zero value
	fmt.Println(m["phone"])

	m2 := make(map[string]int)
	m2["age"] = 22
	m2["price"] = 220
	fmt.Println(m2["age"])
	fmt.Println(m2["phone"])

	fmt.Println(len(m2))

	delete(m2, "price")
	fmt.Println(m2)

	clear(m2)
	fmt.Println(m2)

	m3 := map[string]int{"price": 40, "phones": 3}
	fmt.Println(m3)

	v, ok := m3["priced"]
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	}else {
		fmt.Println("Not ok")
	}

	m4 := map[string]int{"price": 40, "phones": 3}
	m5 := map[string]int{"price": 40, "phones": 6}

	fmt.Println(maps.Equal(m4, m5))

}
