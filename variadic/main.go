package main

import "fmt"

func main() {
	// numbers := []int{1, 29, 23}
	sum := sumup(1, 20, 10)
	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum = sum + val
	}

	return sum
}
