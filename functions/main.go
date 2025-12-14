package main

import "fmt"

type transfromFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 6, 7, 8}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Number: ", numbers)
	fmt.Println("Doubled Number: ", doubled)
	fmt.Println("Tripled Number: ", tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformNumbers(&numbers, transformerFn1)
	transformNumbers(&moreNumbers, transformerFn2)
}

func transformNumbers(number *[]int, transform transfromFn) []int {
	dNumbers := []int{}
	for _, value := range *number {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func getTransformerFunction(number *[]int) transfromFn {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
