package main

import "fmt"

func initial() {
	// Constants declaration
	const pi float64 = 3.14
	const pi2 = 3.14 // inferring type

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Variable declaration
	base := 12          // := means that variable hasn't been created
	var altura int = 70 // with keyword var
	var area int        // without asignation

	// ALL variables or constants must be used, if not, the code will not compile

	println(base, altura, area)

	// Zero values
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBoolean bool

	println(zeroInt, zeroFloat, zeroString, zeroBoolean)
	// 0, 0.0, "", false
	// In go, by default all the variables that don't have an assignation, are zero (numbers), empty strings (string) or false (boolean) by default

	// Area Cuadrado
	const cuadrado = 10
	baseCuadrado := cuadrado * cuadrado
	println("Area Cuadrado: ", baseCuadrado)

	// Valores primitivos
	/**
	Integers
	int = depende del OS (32 or 64 bits)
	int8 = 8 bits = -128 a 127
	int16 = 16 bits
	int32 = 32 bits
	int64 = 64 bits

	Positives Integers
	uint = According to OS (32 or 64 bits)
	uint8 = 8 bits = 0 a 2^8 - 1
	uint16 = 16 bits = 0 a 2^16 - 1
	uint32 = 32 bits = 0 a 2^32 - 1
	uint64 = 64 bits = 0 a 2^64 - 1

	Decimals Numbers
	float32 = 32 bits
	float64 = 64 bits

	Strings and Booleans
	string = ""
	bool = true or false

	Complex Numbers
	complex64 = Real e imaginario float32
	complex128 = Real e imaginario float64
	c := 10 = 8i
	*/
}

func learningFmt() {
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage) // Adds an space

	nombre := "Ivan"
	edad := 25

	// PrintF
	fmt.Printf("%s is %d years old \n", nombre, edad) // %s for string, %d for int
	fmt.Printf("%v tiene %v años \n", nombre, edad)   // %v if we don't know type of data

	// SprintF
	message := fmt.Sprintf("%s is %d years old", nombre, edad) // Sprintf will create a string and it gonna be saved on message
	println(message)

	// Getting data type
	fmt.Printf("helloMessage: %T", helloMessage) // %T will print datatype

}

// Returning a value. a, b int mean that a and b are both int
func multiply(a, b int) int {
	return a * b
}

// We can make functions that return more than one value. This function return text and result, string and int
func square(a int) (text string, result int) {
	return fmt.Sprintf("The square of %d is:", a), a * a
}

func testingReturns() {
	a, b := 3, 4

	multiply := multiply(a, b)
	println(multiply)

	// Getting two returns
	textA, resultA := square(a)
	println(textA, resultA)

	// If we need to omit one of the result, we can do that using "_" for ignoring that result
	_, resultB := square(b)
	println(resultB)
}

// godoc.org for documentation
// Cobra package for create a CLI
