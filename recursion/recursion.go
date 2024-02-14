package main

import "fmt"

func main() {
	number := 6
	fac := factorial(number)
	fmt.Printf("%v!: %v\n", number, fac)

	fmt.Println("-----------------Recursion---------------")
	fac = factorialRecursion(number)
	fmt.Printf("%v!: %v\n", number, fac)
}

// factorial => 5! => 5 * 4 * 3 * 2 * 1 = 120
func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result *= i
	}

	return result
}

func factorialRecursion(number int) int {
	// defined exit condition, otherwise will be in infinity loop
	if number <= 1 {
		return 1
	}
	return number * factorialRecursion(number-1)
}
