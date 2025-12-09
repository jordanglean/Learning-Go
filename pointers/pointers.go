package main

import "fmt"

func main() {
	// regular variable
	age := 32

	// example pointer
	var agePointer *int

	agePointer = &age

	// will output address
	fmt.Println("Age: ", agePointer) // > Age:  0xc00011a008

	// dereference pointer
	fmt.Println("Age: ", *agePointer) // > Age:  32

	adultYears := getAdultYears(&age)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
