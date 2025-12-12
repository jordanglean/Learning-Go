package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"A book"}

	prices := [4]float64{10.99, 9.99, 45.33, 23.00}

	productNames[1] = "A Carpet"

	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	fmt.Println(featuredPrices)
}
