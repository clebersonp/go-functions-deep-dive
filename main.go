package main

import "fmt"

// store function type as custom type
type transformFn func(int) int

// just for example of a more complex func signature as a custom type
type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {

	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	// function as first class values, passing function around as values just as you can pass other kind of values around
	// functions are just values in Go at the end
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	var transformResult []int // nil slice declaration
	for _, number := range *numbers {
		transformResult = append(transformResult, transform(number))
	}
	return transformResult
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
