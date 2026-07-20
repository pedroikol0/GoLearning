package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	//1
	myHobbies := [3]string{"Reading", "Gaming", "Studying Languages"}
	fmt.Println(myHobbies)

	//2
	fmt.Println(myHobbies[0])
	fmt.Println(myHobbies[1:3])

	//3
	myHobbies2 := myHobbies[0:2]
	myHobbies3 := myHobbies[:2]
	fmt.Println(myHobbies2)
	fmt.Println(myHobbies3)

	//4
	fmt.Println(myHobbies3[0:3])

	//5
	myGoals := []string{"Master go, ", "Being a better programmer, "}
	fmt.Println(myGoals)

	//6
	myGoals[1] = "Being a junior go developer"
	myGoals = append(myGoals, "Find a remote go job")
	fmt.Println(myGoals)

	//7
	products := []Product{
		{
			"first-product",
			"A first project",
			12.99,
		},
		{
			"secound-product",
			"A Secound project",
			98.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A Third product",
		27.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
