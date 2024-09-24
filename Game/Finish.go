package game

import (
	"fmt"
	"time"
)

func DisplayFinish() bool {
	if Win {
		ClearScreen()
		fmt.Println(Red, " ============================== Hangman game ==============================") //à continuer
		fmt.Printf("Congratulations\n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Do you want to replay ?\n(1) Yes  (2) No\n")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "1":
			return true
		case "2":
			return false
		default:
			DisplayFinish()
		}
	} else {
		ClearScreen()
		fmt.Println(Red, " ============================== Hangman game ==============================")
		fmt.Println()
		fmt.Println("******************************   Game over   ****************************")
		fmt.Println("|                                                                       |")
		fmt.Println("|                         Nice try, the word was                        |")
		fmt.Printf("|%s|\n", Formatage(Word, 70))
		time.Sleep(2 * time.Second)
		fmt.Printf("The word was : %s.\n\n", Word)
		DisplayHangman(0)
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\nDo you want to retry ?\n(1) Yes  (2) No\n")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "1":
			return true
		case "2":
			return false
		default:
			DisplayFinish()
		}
	}
	return false
}
