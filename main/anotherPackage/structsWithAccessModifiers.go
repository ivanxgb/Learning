package anotherPackage

import "fmt"

// CarPublic : car with access public, this because it starts with uppercase
type CarPublic struct {
	Brand string // starting with uppercase will be public
	Year  int
}

// carPrivate : car with access private, this because it starts with lowercase
type carPrivate struct {
	brand string // starting with lowercase will be private
	year  int
}

func PrintMessage(text string) {
	fmt.Println(text)
}

func PublicFunction() {
	privateFunction()
}

func privateFunction() {
	PrintMessage("This function is been called since PublicFunction, and cannot be accessed out this package")
}
