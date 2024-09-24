package main

import (
	"fmt"
	filereader "main/FileReader"
	game "main/Game"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("You forgot the file name, try with this command : go run . Ficher1.txt")
	} else {
		WordToFinf := filereader.RandomWord(os.Args[1])
		game.Start(WordToFinf)
		Continue := game.DisplayFinish()
		if Continue {
			main()
		}
	}
}
