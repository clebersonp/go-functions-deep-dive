package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumUp(numbers)
	fmt.Println("Without variadic:", sum)

	// when using a variadic, we need to put three dots after variable do spread it into single element
	sum = sumUpVariadic("------------------", numbers...)
	fmt.Println("With variadic using spread operator:", sum)
	sum = sumUpVariadic("------------------", 1, 10, 15)
	fmt.Println("With variadic passing numbers manually:", sum)

}

func sumUp(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// sumUpVariadic is like a varargs in java.
// if function has more than one parameter, so variadic must be the last one, otherwise compile will complain
// at the end variadic will be a slice of parameter type
func sumUpVariadic(msg string, numbers ...int) int {
	fmt.Println(msg)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
