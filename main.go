package main

import (
	"fmt"

	"github.com/msharaf315/Learning-GO/characters"
)

// game variables
var winner string
var difficulty int
var level int
var turn int

//character variables
var player *characters.Player
var goblin *characters.Goblin
var dragon *characters.Dragon

func main() {

	// Start the game
	err := startGame()
	if err != nil {
		fmt.Println("GAME CRASHED WHILE STARTING 💀")
		return
	}
	// Game loop
	for winner == "" {
		// execute game loop
		winner = gameLoop()
	}

	fmt.Println(winner)
}
