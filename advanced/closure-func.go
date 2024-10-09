package main

import "fmt"

// A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

func closureFunctionExample() {
	fmt.Println("***CLOSURE FUNCTION USAGE EXAMPLE***")
	newConcatter := concatter()
	newConcatter("This")
	newConcatter("string")
	newConcatter("has")
	newConcatter("been")
	newConcatter("concattenated")
	newConcatter("and")
	newConcatter("finally")
	fmt.Println(newConcatter("printed."))
}

func concatter() func(string) string {
	accum := ""
	return func(str string) string {
		accum += str + " "
		return accum
	}
}
