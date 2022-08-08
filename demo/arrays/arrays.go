package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [3]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is cleaned!")
		} else {
			fmt.Println(room.name, "needs to be cleaned!")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Dev"},
		{name: "Reception"},
	}

	checkCleanliness(rooms)

	fmt.Println("Performing cleaning...")

	rooms[1].cleaned = true
	rooms[2].cleaned = true

	checkCleanliness(rooms)
}
