package main

import (
	"fmt"

	"example.com/bank/fileops"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error: ")
		fmt.Println(err)
		fmt.Println("==============================")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Bank Phone Number: ", randomdata.PhoneNumber())

	for {

		presentOptions()

		var choice int
		fmt.Print(">> ")
		_, _ = fmt.Scan(&choice)
		fmt.Println("Your choice: ", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			_, _ = fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New ammount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:

			fmt.Print("Your ammount to withdraw: ")
			var withdrawAmmount float64
			_, _ = fmt.Scan(&withdrawAmmount)

			if withdrawAmmount > accountBalance {
				fmt.Println("You can't withdraw more than you have!")
				continue
			}

			accountBalance -= withdrawAmmount
			fmt.Println("Balance updated! New ammount: ", accountBalance)

			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Exited Go Bank")
			return
		}

	}
}
