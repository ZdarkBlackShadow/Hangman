package game

import (
	"fmt"
	"time"
)

func DisplayFinish() bool {
	/*
		Function who take in charge the end of the Game.
	*/
	if Win {
		ClearScreen()
		fmt.Println(Red, " ============================== Hangman game ==============================")
		fmt.Println()
		fmt.Println(Green, " ***************************** Congratulations *****************************")
		DisplayNbGameDifficulty()
		fmt.Println(Yellow, "|                          Do you want to replay ?                       ", Yellow, "|")
		fmt.Println(Yellow, "|                                                                          |")
		fmt.Println(" |", Cyan, "                     (1)Yes                   (2)No                   ", Yellow, "|")
		fmt.Println(Yellow, "|                                                                          |")
		fmt.Println(Green, " *************************************************************************")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "1":
			return ChangeLevelUp()
		case "2":
			return false
		default:
			return DisplayFinish()
		}
	} else {
		ClearScreen()
		fmt.Println(Red, " ============================== Hangman game ==============================")
		fmt.Println()
		fmt.Println("  ******************************   Game over   *****************************")
		DisplayNbGameDifficulty()
		fmt.Println(Yellow, "|", Orange, "                        Nice try, the word was                        ", Yellow, "|")
		fmt.Println(Yellow, "|                                                                          |")
		Formatage()
		fmt.Println(Yellow, "|                                                                          |")
		DisplayHangman2(0)
		fmt.Println(Yellow, "|                                                                          |")
		fmt.Println(Yellow, "|                          Do you want to replay ?                       ", Yellow, "|")
		fmt.Println(Yellow, "|                                                                          |")
		fmt.Println(" |", Cyan, "                    (1)Yes                   (2)No                    ", Yellow, "|")
		fmt.Println(Yellow, "|                                                                          |")
		fmt.Println(Red, " **************************************************************************")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "1":
			return ChangeLevelDown()
		case "2":
			return false
		default:
			return DisplayFinish()
		}
	}
}

func ChangeLevelUp() bool {
	//Function who propose to increase the difficulty
	ClearScreen()
	fmt.Println(Red, " ============================== Hangman game ==============================")
	fmt.Println()
	fmt.Println(Green, " **************************************************************************")
	DisplayNbGameDifficulty()
	fmt.Println(Yellow, "|                ", Orange, "do you want to increase the difficulty ?", Yellow, "              |")
	fmt.Println(Yellow, "|                                                                          |")
	fmt.Println(" |", Red, "                  (1)Yes", Green, "                   (2)No                    ", Yellow, "|")
	fmt.Println(Yellow, "|                                                                          |")
	fmt.Println(Green, " **************************************************************************")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		if Difficulte != 10 {
			Difficulte += 1
		}
		if Filename == "File10.txt" {
			ClearScreen()
			fmt.Println(Red, " ============================== Hangman game ==============================")
			fmt.Println()
			fmt.Println(Green, " **************************************************************************")
			DisplayNbGameDifficulty()
			fmt.Println(Yellow, "|                You're already at maximum difficuly level.                |")
			fmt.Println(Yellow, "|                                                                          |")
			fmt.Println(Yellow, "|                   You will continue in a short while.                    |")
			fmt.Println(Yellow, "|                                                                          |")
			fmt.Println(Green, " **************************************************************************")
			time.Sleep(3 * time.Second)
			return true
		} else {
			switch Filename[4] {
			case '1':
				if Filename[5] == '0' {
					Filename = "File10.txt"
				} else {
					Filename = "File2.txt"
				}
			case '2':
				Filename = "File3.txt"
			case '3':
				Filename = "File4.txt"
			case '4':
				Filename = "File5.txt"
			case '5':
				Filename = "File6.txt"
			case '6':
				Filename = "File7.txt"
			case '7':
				Filename = "File8.txt"
			case '8':
				Filename = "File9.txt"
			}
			return true
		}
	case "2":
		return true
	default:
		return ChangeLevelUp()
	}
}

func ChangeLevelDown() bool {
	//Function who propose to decrease the difficulty
	ClearScreen()
	fmt.Println(Red, " ============================== Hangman game ==============================")
	fmt.Println()
	fmt.Println(Green, " *************************************************************************")
	DisplayNbGameDifficulty()
	fmt.Println(Yellow, "|               ", Orange, "do you want to decrease the difficulty ?", Yellow, "               |")
	fmt.Println(Yellow, "|                                                                          |")
	fmt.Println(" |", Red, "                  (1)Yes", Green, "                   (2)No                    ", Yellow, "|")
	fmt.Println(Yellow, "|                                                                          |")
	fmt.Println(Green, " *************************************************************************")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		if Difficulte != 1 {
			Difficulte -= 1
		}
		if Filename == "File1.txt" {
			ClearScreen()
			fmt.Println(Red, " ============================== Hangman game ==============================")
			fmt.Println()
			fmt.Println(Green, " **************************************************************************")
			DisplayNbGameDifficulty()
			fmt.Println(Yellow, "|                You're already at minimum difficuly level.                |")
			fmt.Println(Yellow, "|                                                                          |")
			fmt.Println(Yellow, "|                   You will continue in a short while.                    |")
			fmt.Println(Yellow, "|                                                                          |")
			fmt.Println(Green, " **************************************************************************")
			time.Sleep(3 * time.Second)
			return true
		} else {
			switch Filename[4] {
			case '1':
				Filename = "File9.txt"
			case '9':
				Filename = "File8.txt"
			case '8':
				Filename = "File7.txt"
			case '7':
				Filename = "File6.txt"
			case '6':
				Filename = "File5.txt"
			case '5':
				Filename = "File4.txt"
			case '4':
				Filename = "File3.txt"
			case '3':
				Filename = "File2.txt"
			case '2':
				Filename = "File1.txt"
			}

			return true
		}
	case "2":
		return true
	default:
		return ChangeLevelDown()
	}
}
