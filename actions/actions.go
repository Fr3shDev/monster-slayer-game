package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = 100
var currentPlayerHealth = 100

func AttackMonster(isSpecialAttack bool) {
	minAttackValue := 5
	maxAttackValue := 10

	if isSpecialAttack {
		minAttackValue = 10
		maxAttackValue = 20
	}
	damageValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth = currentMonsterHealth - damageValue

}

func HealPlayer() {
	minHealValue := 10
	maxHealValue := 20

	healValue := generateRandBetween(minHealValue, maxHealValue)

	healthDiff := 100 - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = 100
	}
}

func AttackPlayer() {
	minAttackValue := 9
	maxAttackValue := 13

	damageValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentPlayerHealth -= damageValue
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
