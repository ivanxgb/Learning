package anotherPackage

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

// GetBrand : This will add the function as part of the struct, something like extension functions
func (myPc Pc) GetBrand() {
	fmt.Println(myPc.Brand) // Printing Brand
}

func (myPc Pc) GetRam() {
	fmt.Println(myPc.Ram) // Printing Ram
}

// DuplicateRam will take a reference to myPc object and update its value
func (myPc *Pc) DuplicateRam() {
	myPc.Ram *= 2
}

// String will override the function that returns a value when printing
func (myPc Pc) String() string {
	return fmt.Sprintf("Pc of %d RAM, %d Disk from %s", myPc.Ram, myPc.Disk, myPc.Brand)
}
