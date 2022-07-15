package main

import "fmt"

// Learning about cycles, for, for while, for forever
func testingCycles() {
	// learningFor()
	// learningForWhile()
	// learningForForever()
	learningForRange()
}

func learningFor() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func learningForWhile() {
	counter := 0
	for counter < 10 {
		fmt.Printf("counter: %d \n", counter)
		counter++
	}
}

func learningForForever() {
	counter := 0
	for {
		fmt.Printf("Counter forever %d \n", counter)
		counter++
	}
}

func learningForRange() {
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}
}
