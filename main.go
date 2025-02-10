package main

import "github.com/Fr3shDev/monster-slayer-game/interaction"

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {}

func endGame() {}
