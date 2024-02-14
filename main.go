package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	doubleTransformed := transformNumbers(&numbers, func(number int) int { return number * 2 })
	fmt.Println(doubleTransformed)

	tripledTransformed := transformNumbers(&numbers, func(number int) int { return number * 3 })
	fmt.Println(tripledTransformed)

	fmt.Println("--------------------")
	// factory function pattern
	// closure: will lock in the factors 2 and 3 inside an anonymous function
	double := createTransformer(2)
	triple := createTransformer(3)

	// passing functions as values
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	var dNumbers []int

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// createTransformer is a factory function pattern to create functions
func createTransformer(factor int) func(int) int {
	// every anonymous functions are closures in Go and lock in values of the out of scope (factor in this case)
	// for future functions execution
	return func(number int) int {
		return number * factor
	}
}
