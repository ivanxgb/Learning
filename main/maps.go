package main

import "fmt"

func testingMaps() {
	learningMaps()
}

func learningMaps() {
	m := make(map[string]int) // map[key]value

	m["ivan"] = 24
	m["stef"] = 24

	fmt.Println(m)

	// for traversing we will use a for range, instead of index we will have the key
	for key, val := range m {
		fmt.Println(key, val)
	}

	// for accessing values
	fmt.Println(m["ivan"])

	// If we try access to a value that hasn't been declared, it will throw a zero value
	// In order to solve this, we can use another variables like destructuring

	// value, ok := map["key"]. If ok is false, that will mean that key doesn't exist

	value, ok := m["nobody"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Doesn't exist") // Will print this bc' ok is false
	}

	value, ok = m["stef"]
	if ok {
		fmt.Println(value)
	}
}
