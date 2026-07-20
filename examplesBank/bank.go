package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----------------")
		return
		//panic("Can't continue, sorry!")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("Welcome to Go Bank")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdrown money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated: ", accountBalance)
		default:
			fmt.Println("Goodbye!")
			return
		}
	}
}
