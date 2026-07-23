package main

import (
	"fmt"

	"bankTest.go/fileInteractions"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileInteractions.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----------------")
		return
		//panic("Can't continue, sorry!")
	}

	fmt.Println("Welcome to Go Bank!")

	for {

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("How much do you want to deposit: ")
			var newDeposit float64
			fmt.Scan(&newDeposit)

			if newDeposit <= 0 {
				fmt.Println("Please insert a valid number!")
				continue
			}

			accountBalance += newDeposit
			fileInteractions.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("Balance updated: ", accountBalance)
		case 3:
			fmt.Print("How much you need to withdrown: ")
			var newWithdrown float64
			fmt.Scan(&newWithdrown)

			if newWithdrown <= 0 {
				fmt.Println("Please insert a valid number!")
				continue
			}

			if newWithdrown > accountBalance {
				fmt.Println("You can't withdrown more than you have in the account.")
				continue
			}

			accountBalance -= newWithdrown
			fileInteractions.WriteFloatToFile(accountBalanceFile, accountBalance)
			fmt.Println("Balance updated: ", accountBalance)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}
