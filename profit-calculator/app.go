package main

import "fmt"

const taxRate = 0.20

func main() {
	var revenue float64
	var expenses float64

	revenue = getUserInput("Enter the ammount of revenue: ")

	expenses = getUserInput("Enter the ammount of expenses: ")

	ebt, profit, ratio := calculateEarnings(revenue, expenses)

	fmt.Print("Earnings before tax: ")
	fmt.Println(ebt)

	fmt.Print("Ratio: ")
	fmt.Println(ratio)

	fmt.Print("Profit: ")
	fmt.Println(profit)
}

func getUserInput(text string) float64 {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}

func calculateEarnings(revenue float64, expenses float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate)
	ratio = ebt * (1 - taxRate)

	return
}
