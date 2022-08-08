//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

//* Create a function to print out the contents of the assembly line
func printContent(assembly []Part) {
	fmt.Println("\nAssembly line content:")
	for i := 0; i < len(assembly); i++ {
		item := assembly[i]
		fmt.Println(i, item)
	}
}

func main() {

	//* Using a slice, create an assembly line that contains type Part
	assembly := make([]Part, 3)

	//  - Create an assembly line having any three parts
	assembly[0] = "GPU"
	assembly[1] = "CPU"
	assembly[2] = "Motherboard"
	printContent(assembly)

	//  - Add two new parts to the line
	assembly = append(assembly, "HDD", "SDD")
	printContent(assembly)

	//  - Slice the assembly line so it contains only the two new parts
	assembly = assembly[3:]
	printContent(assembly)

}
