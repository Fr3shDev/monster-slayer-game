package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action              string
	PlayerAttackDamage  int
	PlayerHealValue     int
	MonsterAttackDamage int
	PlayerHealth        int
	MonsterHealth       int
}

func PrintGreeting() {
	asciiFigure := figure.NewColorFigure("MONSTER SLAYER", "", "green", true)
	asciiFigure.Print()
	fmt.Println("Starting a new game...")
	fmt.Println("Good Luck!!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")
 
	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDamage)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage.\n", roundData.PlayerAttackDamage)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDamage)
			fmt.Printf("Player Health %v \n", roundData.PlayerHealth)
			fmt.Printf("Monster Health %v \n", roundData.MonsterHealth)


}

func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	asciiFigure := figure.NewColorFigure("GAME OVER", "", "red", true)
	asciiFigure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {
	exPath, error := os.Executable()

	if error != nil {
		fmt.Println("Writing log file failed. Exiting!")
		return
	}

	exPath = filepath.Dir(exPath)

	file, error := os.Create(exPath + "/gamelog.txt")
	// file, error := os.Create("gamelog.txt") // For "go run"

	if error != nil {
		fmt.Println("Saving a log file failed. Exiting.")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round": fmt.Sprint(index + 1),
			"Action": value.Action,
			"Player Attack Damage": fmt.Sprint(value.PlayerAttackDamage),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDamage),
			"Player Health": fmt.Sprint(value.PlayerHealth),
			"Monster Health": fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, error = file.WriteString(logLine)

		if error != nil {
			fmt.Println("Writing into log file failed. Exiting.")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log!")
}
