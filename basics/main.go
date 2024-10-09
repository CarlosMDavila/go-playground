package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	pointerExample()
	fmt.Print("\n\n")
	ifScopedVariableExample()
	fmt.Print("\n\n")
	namedReturnsExample()
	fmt.Print("\n\n")
	embeddedStructExample()
	fmt.Print("\n\n")
	interfaceExample()
	fmt.Print("\n\n")
	whileExample()
	fmt.Print("\n\n")
	variadicFunctionExample()
	fmt.Print("\n\n")
	structMapKeyExample()
}

func pointerExample() {
	fmt.Println("***POINTER USAGE EXAMPLE***")
	var originalVar int = 42
	var copiedValue int = originalVar
	var pointerVar *int = &originalVar
	fmt.Println("Original value: ", originalVar)
	fmt.Println("Copied value: ", copiedValue)
	fmt.Println("Pointer value: ", *pointerVar)

	originalVar = 45
	fmt.Println("-Original value changed-")
	fmt.Println("Copied value: ", copiedValue)
	fmt.Println("Pointer value: ", *pointerVar)

	*pointerVar = 64
	fmt.Println("-Pointer referenced value changed-")
	fmt.Println("Original value: ", originalVar)
	fmt.Println("Pointer value: ", *pointerVar)
}

func ifScopedVariableExample() {
	fmt.Println("***IF SCOPED VARIABLE EXAMPLE***")
	if value := rand.IntN(100); value > 50 {
		fmt.Println("Random value is greater than 50")
	} else {
		fmt.Println("Random value is ", value)
	}
}

func namedReturnsExample() {
	fmt.Println("***NAMED RETURN EXAMPLE***")
	a, b := rand.IntN(100), rand.IntN(100)

	sum, sub, mult, div, err := basicOperations(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Num 1: %d, Num 2: %d. Sum %d, Substraction: %d, Multiplication: %d, Division: %d.\n", a, b, sum, sub, mult, div)
}

func basicOperations(a, b int) (sum, substraction, multiplication, division int, err error) {
	if b == 0 {
		return 0, 0, 0, 0, errors.New("can't divide by zero")
	}
	sum = a + b
	substraction = a - b
	multiplication = a * b
	division = a / b
	return sum, substraction, multiplication, division, nil
}

func embeddedStructExample() {
	fmt.Println("***EMBEDDED STRUCT USAGE EXAMPLE***")
	truck := Truck{
		bedSize: 5,
		Car: Car{
			maker: "Toyota",
			model: "X5",
		},
	}

	fmt.Println("Truck's bed size: ", truck.bedSize)
	fmt.Println("Truck's maker: ", truck.maker)
	fmt.Println("Truck's model: ", truck.model)
}

func interfaceExample() {
	fmt.Println("***INTERFACE USAGE EXAMPLE***")
	shapeList := []Shape{}
	shapeList = append(shapeList, Circle{radius: 5})
	shapeList = append(shapeList, Square{side: 10})

	for _, shape := range shapeList {
		_, ok := shape.(Circle)
		if !ok {
			fmt.Print("This one is not a circle. ")
		}
		switch s := shape.(type) {
		case Circle:
			fmt.Println("The area of the circle is:", s.Area())
		case Square:
			fmt.Println("The area of the square is:", s.Area())
		}
	}
}

func whileExample() {
	fmt.Println("***WHILE (FOR SIMULATION) USAGE EXAMPLE***")
	countDown := 5
	for countDown > 0 {
		fmt.Println("Ticks left: ", countDown)
		time.Sleep(500 * time.Millisecond)
		countDown--
	}
	fmt.Println("Timer has reached 0.")
}

func variadicFunctionExample() {
	fmt.Println("***VARIADIC USAGE EXAMPLE***")
	printStrings("THIS", "ARE", "PARAMETERED", "VALUES")

	stringSlice := []string{"THIS", "IS", "A", "SLICE"}
	printStrings(stringSlice...)
}

func printStrings(strings ...string) {
	for _, str := range strings {
		fmt.Println("Input word:", str)
	}
}

func structMapKeyExample() {
	fmt.Println("***STRUCT MAP KEY USAGE EXAMPLE***")

	carSales := make(map[Car]int)
	carSales[Car{maker: "Toyota", model: "Camry"}] = 10
	carSales[Car{maker: "Mazda", model: "CX-5"}] = 20
	carSales[Car{maker: "Toyota", model: "Prius"}] = 30

	for key, value := range carSales {
		fmt.Println(key.maker, "sold", value, key.model)
	}
}
