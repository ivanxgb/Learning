package main

// Learning to use keywords defer, break y continue

func testingKeywords() {
	learningDefer()
	learningContinue()
	learningBreak()
}

// Defer will defer a statement to the end. Use case: close connection to a db, close a file in order to drop something for memory
func learningDefer() {
	defer println("Testing defer. This will print last")
	println("This will print first")
}

// Will skip the cycle to next execution
func learningContinue() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}
}

// Gonna break the execution
func learningBreak() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}
}
