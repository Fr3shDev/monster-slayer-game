package main

import (
	"github.com/Fr3shDev/monster-slayer-game/actions"
	"github.com/Fr3shDev/monster-slayer-game/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	
	var playerAttackDamage int
	var playerHealValue int
	var monsterAttackDamage int 

	if userChoice == "ATTACK" {
		playerAttackDamage = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDamage = actions.AttackMonster(true)
	}

	monsterAttackDamage = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action: userChoice,
		PlayerHealth: playerHealth,
		MonsterHealth: monsterHealth,
		PlayerAttackDamage: playerAttackDamage,
		PlayerHealValue: playerHealValue,
		MonsterAttackDamage: monsterAttackDamage,
	}
	interaction.PrintRoundStatistics(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
