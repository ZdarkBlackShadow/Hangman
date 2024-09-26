package main

import (
	"fmt"
	filereader "main/FileReader"
	game "main/Game"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("You forgot the file name, try with this command : go run . File1.txt")
	} else {
		if !game.InGame {
			game.Filename = os.Args[1]
		}
		WordToFinf := filereader.RandomWord(game.Filename)
		game.Start(WordToFinf)
		Continue := game.DisplayFinish()
		if Continue {
			main()
		}
	}
}
