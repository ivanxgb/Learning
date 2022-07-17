package main

import (
	ap "Learning/main/learningStructs"
	"fmt"
)

func testingStructs() {
	// learningLocalStruct()
	// learningStructInOtherPackage()
	// structsAndPointers()
	// learningStructsWithPointers()
	learningToCustomizeOutput()
}

func learningLocalStruct() {
	// Creating an instance
	myCar := Car{"Ford", 2020}
	fmt.Println(myCar) // {Ford 2020}

	// Another way to do...
	var otherCar Car
	otherCar.brand = "Chevrolet"
	// If we don't declare model, it will take a zero value
	fmt.Println(otherCar) // {Chevrolet 0}
}

/* Structs are the way to declare classes */

type Car struct {
	brand string
	model int
}

func learningStructInOtherPackage() {

	var myCar = ap.CarPublic{ // in order to have access to a struct in a different package, we must write entire directory
		Brand: "ivan",
		Year:  2020,
	}
	fmt.Println(myCar)

	ap.PrintMessage("hola")
	ap.PublicFunction()
}

func structsAndPointers() {
	a := 50 // This will save a value
	b := &a // This will save a pointer to a memory space

	fmt.Println(a)  // 50
	fmt.Println(b)  // 0x0----- (address in memory)
	fmt.Println(*b) // 50

	*b = 10        // Bc *b is a pointer, 'a' will take the new value
	fmt.Println(a) // 10
}

func learningStructsWithPointers() {
	myPc := ap.Pc{
		Ram:   16,
		Disk:  1024,
		Brand: "MSI",
	}

	fmt.Println(myPc)
	myPc.GetBrand()

	myPc.DuplicateRam() // This will duplicate ram
	myPc.GetRam()

	myPc.DuplicateRam() // This will duplicate ram again
	myPc.GetRam()
}

func learningToCustomizeOutput() {
	myPc := ap.Pc{
		Ram:   16,
		Disk:  32,
		Brand: "Apple",
	}

	fmt.Println(myPc) // "Pc of 16 RAM, 32 Disk from Apple" will be printed bc the function was overridden
}
