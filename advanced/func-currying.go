package main

import "fmt"

// Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function.

func functionCurryingExample() {
	fmt.Println("**FUNCTION CURRYNG USAGE EXAMPLE***")
	squareFunction := selfMath(multiplyInt)
	doubleFunction := selfMath(addInts)

	fmt.Println("Square function for val 8:", squareFunction(8))
	fmt.Println("Double function for val 16:", doubleFunction(16))
}

func selfMath(arithmeticOperation func(int, int) int) func(int) int {
	return func(val1 int) int {
		return arithmeticOperation(val1, val1)
	}
}
