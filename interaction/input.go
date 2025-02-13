package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	PlayerChoiceAttack        = 1
	PlayerChoiceHeal          = 2
	PlayerChoiceSpecialAttack = 3
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAttackIsAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == fmt.Sprint(PlayerChoiceAttack) {
			return "ATTACK"
		} else if playerChoice == fmt.Sprint(PlayerChoiceHeal) {
			return "HEAL"
		} else if playerChoice == fmt.Sprint(PlayerChoiceSpecialAttack) && specialAttackIsAvailable {
			return "SPECIAL_ATTACK"
		}

		fmt.Println("Fetching the user input failed. Please try again.")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	userInput, error := reader.ReadString('\n')

	if error != nil {
		return "", error
	}

	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput, nil
}
