package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	doubleTransformed := transformNumbers(&numbers, func(number int) int { return number * 2 })
	fmt.Println(doubleTransformed)

	tripledTransformed := transformNumbers(&numbers, func(number int) int { return number * 3 })
	fmt.Println(tripledTransformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	var dNumbers []int

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
