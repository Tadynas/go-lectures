//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greeting(name string) {
	fmt.Println("Greetings", name)
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func message() string {
	return "Hello world!"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func add(x, y, z int) int {
	return x + y + z
}

//* Write a function that returns any number
func number() int {
	return rand.Intn(100)
}

// //* Write a function that returns any two numbers
func twoNumbers() (int, int) {
	return number(), number()
}

//* Add three numbers together using any combination of the existing functions.
//  * Print the result
func addThreeNumbers() {
	a := number()
	b, c := twoNumbers()
	fmt.Println("Add three random numbers", add(a, b, c))
}

//* Call every function at least once
func main() {
	greeting("Ted")

	fmt.Println(message())

	fmt.Println("Sum of 3 + 5 + 2 is:", add(3, 5, 2))

	fmt.Println("Random number:", number())

	a, b := twoNumbers()

	fmt.Println("Two random numbers:", a, b)

	addThreeNumbers()
}
