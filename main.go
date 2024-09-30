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
		temp1, temp2 := filereader.RandomWord(game.Filename)
		game.Difficulte = temp2
		game.Start(temp1)
		Continue := game.DisplayFinish()
		if Continue {
			main()
		}
	}
}
