package main

import "fmt"

type transfromFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Number: ", numbers)
	fmt.Println("Doubled Number: ", doubled)
	fmt.Println("Tripled Number: ", tripled)
}

func transformNumbers(number *[]int, transform transfromFn) []int {
	dNumbers := []int{}
	for _, value := range *number {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
