//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
type SecurityTag struct {
	state bool
}

//* Create functions to activate and deactivate security tags using pointers
func changeState(securityTag *SecurityTag, state bool) {
	securityTag.state = state
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(securityTags []SecurityTag) {
	for i := 0; i < len(securityTags); i++ {
		changeState(&securityTags[i], false)
	}
}

func main() {

	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	securityTags := []SecurityTag{
		{state: true},
		{state: true},
		{state: true},
		{state: true},
	}
	fmt.Println("\n1. All activated")
	fmt.Println(securityTags)

	//  - Deactivate any one security tag in the array/slice
	changeState(&securityTags[0], false)
	fmt.Println("\n2. First deactivated")
	fmt.Println(securityTags)

	//  - Call the checkout() function to deactivate all tags
	checkout(securityTags)
	fmt.Println("\n3. All deactivated")
	fmt.Println(securityTags)

}
