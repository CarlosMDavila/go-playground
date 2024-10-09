package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Defer allows a function to be executed automatically just before its enclosing function returns.

func deferKeyworkExample() {
	fmt.Println("***DEFER KEYWORD USAGE EXAMPLE***")
	start := time.Now()
	defer func() {
		fmt.Println("Time slept:", time.Since(start))
	}()

	randomInt := rand.IntN(5000)
	time.Sleep(time.Duration(randomInt) * time.Millisecond)
}
