package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (c Salad) PrepareDish() {
	fmt.Println("chop salad")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Caesars Salad")}
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}
