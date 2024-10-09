package main

import "fmt"

// A function can be passed as an argument to other functions, can be returned by another function, and can be assigned as a value to a variable.

func highOrderFunction() {
	fmt.Println("***HIGH ORDER FUNCTION USAGE EXAMPLE***")
	fmt.Println("Sum between 15, 16, 90:", aggregateOperations(15, 16, 90, addInts))
	fmt.Println("Multiplication between 2, 8, 15:", aggregateOperations(2, 8, 15, multiplyInt))
}

func addInts(val1, val2 int) int {
	return val1 + val2
}

func multiplyInt(val1, val2 int) int {
	return val1 * val2
}

func aggregateOperations(val1, val2, val3 int, arithmeticOperation func(int, int) int) int {
	return arithmeticOperation(arithmeticOperation(val1, val2), val3)
}
