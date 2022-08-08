//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health can not go above their maximums
//  - Health can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}
}

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.addHealth(player.maxHealth + 1)
	if player.health > player.maxHealth {
		t.Fatalf("Heal went over limit: %v, want %v", player.health, player.maxHealth)
	}

	if player.health > player.maxHealth {
		t.Fatalf("Health: %v. Maximum: %v", player.health, player.maxHealth)
	}
}
