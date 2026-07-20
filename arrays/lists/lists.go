package main

import "fmt"

func main() {
	prices := []float64{10.99, 29.99, 49.99, 76.99}
	fmt.Println(prices[1])
	prices = append(prices, 89.99) // prices[4] = 89.99  do not work
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"Apple"}

// 	prices := [4]float64{1.99, 13.99, 23.92, 4.99}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	featurePrice := prices[1:3]
// 	fmt.Println(featurePrice)

// }
