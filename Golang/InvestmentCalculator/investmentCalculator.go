package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 1.5

	var investmentAmount, expectedReturnRate, years float64

	fmt.Println("What is the amount of money do you want to invest?")
	fmt.Scan(&investmentAmount)
	fmt.Println("What is the expected return rate?")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("How many years do you want to invest for?")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
