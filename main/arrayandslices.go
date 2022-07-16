package main

import (
	"fmt"
	"strings"
)

func testingArrayAndSlices() {
	//learningArrays()
	//learningSlices()
	//traversingSlices()
	isPalindromeIgnoringCase("amor a romA")
}

// Arrays are immutable, slice are mutable

func learningSlices() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice, len(slice), cap(slice)) // len will return length of slice, cap will return total capacity

	// Slicing...
	println(slice[0])       // will print element at index 0
	fmt.Println(slice[:3])  // will print element until index 3
	fmt.Println(slice[2:4]) // will print elements since index 2 until 4
	fmt.Println(slice[4:])  // will print elements since index 4

	// Append
	slice = append(slice, 25)
	fmt.Println(slice)

	// Append a list to existing slice
	toAppend := []int{26, 27, 28}
	slice = append(slice, toAppend...) // three dots will work as destructuring operator
}

func learningArrays() {
	var ivan [4]string
	ivan[0] = "ivan"
	ivan[1] = "andres"
	ivan[2] = "gonzalez"
	ivan[3] = "barraza"

	stefanny := [4]string{"stefanny", "andrea", "guzman", "barcelo"}

	var fullName string
	for _, name := range ivan {
		fullName += name + " "
	}
	println(fullName)

	fullName = ""
	for _, name := range stefanny {
		fullName += name + " "
	}
	println(fullName)
}

func traversingSlices() {
	slice := []string{"Hi!", "I'm learning", "golang"}

	// index could be ignored using _
	for i, element := range slice {
		fmt.Printf("At index %d: %s \n", i, element)
	}
}

func isPalindrome(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	fmt.Printf("Is palindrome? %v", textReverse == text)
}

func isPalindromeIgnoringCase(text string) {
	var textReverse string
	aux := strings.ToLower(text)
	for i := len(aux) - 1; i >= 0; i-- {
		textReverse += string(aux[i])
	}

	fmt.Printf("Is palindrome? %v", textReverse == aux)
}
