package main

import (
	"log"
	"strconv"
)

func testingConditionals() {
	learningIf()
	learningSwitch(4)
	learningSwitch(3)
}

func learningIf() {
	valor1, valor2 := 4, 5

	println(valor1 == valor2)

	if valor1 >= 3 {
		println("Higher")
	} else {
		println("lower")
	}

	// Convert text to number
	textToNumber("25")
}

func textToNumber(str string) {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln("Cannot convert to string", err)
	}
	println("Value:", val)
}

func learningSwitch(number int) {
	switch number % 2 {
	case 0:
		println("es par")
		break
	default:
		println("es impar")
	}

	switch {
	case number > 10:
		println("Es mayor a 10")
	default:
		println("Es menor a 10")
	}
}
