//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

//  - Products must include the price and the name
type Product struct {
	name  string
	price int
}

func printInfo(shoppingList [4]Product) {
	var cost, totalItems int

	for i := 0; i < len(shoppingList); i++ {
		item := shoppingList[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item on the list:", shoppingList[totalItems-1])
	fmt.Println("Total items:", totalItems)
	fmt.Println("Total cost:", cost)
}

func main() {
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	shoppingList := [4]Product{
		{name: "Banana", price: 1},
		{name: "Coffee", price: 6},
		{name: "Milk", price: 3},
	}

	//* Print to the terminal:
	printInfo(shoppingList)

	//* Add a fourth product to the list and print out the
	//  information again
	shoppingList[3] = Product{name: "Bread", price: 5}
	printInfo(shoppingList)
}
