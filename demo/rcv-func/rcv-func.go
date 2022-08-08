package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(parkingLot *ParkingLot, spaceNum int) {
	parkingLot.spaces[spaceNum-1].occupied = true
}

func (parkingLot *ParkingLot) occupySpace(spaceNum int) {
	parkingLot.spaces[spaceNum-1].occupied = true
}

func (parkingLot *ParkingLot) vacateSpace(spaceNum int) {
	parkingLot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 3)}
	fmt.Println("Initial:", lot)

	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupied:", lot)

	lot.vacateSpace(2)
	fmt.Println("After vacate:", lot)
}
