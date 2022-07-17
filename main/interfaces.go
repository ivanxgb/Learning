package main

import (
	li "Learning/main/learningInterfaces"
	"fmt"
)

func testingInterfaces() {
	// learningInterfaces()
	listOfInterfaces()
}

func learningInterfaces() {
	li.LearningInterfaces()
}

func listOfInterfaces() {

	classicSlide := []int{1, 4, 4} // Classic slice
	fmt.Println(classicSlide)

	myInterface := []interface{}{"Hola", 12, 4.90, true} // Here we're declaring a slice as a list of interface and using two {}{} in order to
	// allow us to accept different types of data
	fmt.Println(myInterface)
}
