package game

import (
	"fmt"
	"math/rand"
	"time"
)

var Finish bool
var Win bool

func Game() {
	rand.Seed(time.Now().UnixNano())
	Finish = false
	Win = false
	ListToDisplay := []rune{}
	AlreadyTry := []string{}
	LetterFind := 0
	RandomLetter := rune(Word[rand.Intn(Lenght)])
	for i := 0; i < Lenght; i++ {
		if rune(Word[i]) == RandomLetter {
			ListToDisplay = append(ListToDisplay, RandomLetter)
			LetterFind++
		} else {
			ListToDisplay = append(ListToDisplay, '_')
		}
	}
	for !Finish {
		ClearScreen()
		fmt.Printf("==============================Hangman game =============================\n\nWord to Find : ")
		for _, element := range ListToDisplay {
			fmt.Printf("%s ", string(element))
		}
		fmt.Printf("\n\nYou already try : ")
		for _, element := range AlreadyTry {
			fmt.Printf("%s ", element)
		}
		fmt.Printf("\n\nTry remaining : %d\n", Lap)
		DisplayHangman(Lap)
		fmt.Printf("\n\nEnter a letter or a word : ")
		var choice string
		fmt.Scan(&choice)
		if len(choice) == 1 {
			SubmitesLetter := choice[0]
			FindInWord := false
			for c := 0; c < Lenght; c++ {
				if SubmitesLetter == Word[c] {
					ListToDisplay[c] = rune(SubmitesLetter)
					FindInWord = true
					LetterFind++
				}
			}
			if !FindInWord {
				AlreadyTry = append(AlreadyTry, choice)
				Lap--
			}
		} else {
			if choice == Word {
				LetterFind = Lenght
			} else {
				AlreadyTry = append(AlreadyTry, choice)
				Lap -= 2
			}
		}
		if LetterFind == Lenght {
			Finish = true
			Win = true
		} else {
			if Lap <= 0 {
				Finish = true
			}
		}
	}
}
