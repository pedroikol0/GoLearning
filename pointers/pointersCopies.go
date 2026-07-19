// package main
// Passing pointers as function arguments

// import "fmt"

// func main() {
// 	age := 32 //regular variable

// 	agePointer := &age //pointer variable

// 	fmt.Println("Age: ", *agePointer)
// 	fmt.Println("Adult years: ", getAdultYears(agePointer))
// }

// func getAdultYears(age *int) int {
// 	return *age - 18
// }

package main

import "fmt"

func main() {
	age := 32 //regular variable

	agePointer := &age //pointer variable

	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)

}

func getAdultYears(age *int) {
	*age = *age - 18
}
