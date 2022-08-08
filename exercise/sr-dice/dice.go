//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dice := 2
	sides := 6
	rolls := 2

	for i := 0; i < rolls; i++ {
		totalRoll := 0

		for j := 0; j < dice; j++ {
			totalRoll += rand.Intn(sides) + 1
		}

		snakeEyes := totalRoll == 2 && dice == 2
		even := totalRoll%2 == 0

		if snakeEyes {
			fmt.Println("Snake Eyes", totalRoll)
		} else if even {
			fmt.Println("Even", totalRoll)
		} else {
			fmt.Println("Odd", totalRoll)
		}

	}
}
