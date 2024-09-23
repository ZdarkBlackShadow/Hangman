package main

import (
	"fmt"
	filereader "main/FileReader"
	game "main/Game"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("You forgot the file name, try with this command : go run . Ficher1.txt")
	} else {
		WordToFinf := filereader.RandomWord(os.Args[1])
		game.Start(WordToFinf)
		DisplayFinish()
	}
}

func DisplayFinish() {
	if game.Win {
		game.ClearScreen()
		fmt.Printf("Congratulations\n")
		time.Sleep(2 * time.Second)
		fmt.Printf("Do you want to replay ?\n(1) Yes  (2) No\n")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "1":
			main()
		case "2":
			return
		default:
			DisplayFinish()
		}
	} else {
		game.ClearScreen()
		fmt.Printf("Nice try.\n\n")
		time.Sleep(2 * time.Second)
		fmt.Printf("The word was : %s.\n\n", game.Word)
		game.DisplayHangman(0)
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\nDo you want to retry ?\n(1) Yes  (2) No\n")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "1":
			main()
		case "2":
			return
		default:
			DisplayFinish()
		}
	}
}
