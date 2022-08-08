//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import (
	"fmt"
)

//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	updatedHealth := player.health + amount
	if updatedHealth <= player.maxHealth {
		player.health = updatedHealth
		fmt.Println(player.name, "Add", amount, "health ->", player.health)
	}
}

func (player *Player) applyDamage(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "Damage", amount, "health ->", player.health)
}

func (player *Player) addEnergy(amount uint) {
	updatedEnergy := player.energy + amount
	if updatedEnergy <= player.maxEnergy {
		player.energy = updatedEnergy
		fmt.Println(player.name, "Add", amount, "energy ->", player.energy)
	}
}

func (player *Player) consumeEnergy(amount uint) {
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Consume energy", amount, "energy ->", player.energy)
}

func main() {
	player := Player{
		name:      "knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	player.applyDamage(99)
	player.addHealth(10)
	player.consumeEnergy(20)
	player.addEnergy(10)

	player.consumeEnergy(9999)
}
