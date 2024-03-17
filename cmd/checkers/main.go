package main

import (
	"checkers/internal/game"
	"fmt"
)

func main() {
	fmt.Println("*********** Start game ***********")
	fmt.Println()
	gameInstance := game.NewGame("player1", "player2")
	fmt.Println(gameInstance)
}
