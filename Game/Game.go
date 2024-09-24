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
		fmt.Println(Red, " ============================== Hangman game ==============================")
		fmt.Println(Yellow, "|                                                                          |")
		temp := ""
		for _, element := range ListToDisplay {
			temp += string(element) + " "
		}
		temp2 := ""
		for i := 0; i < 31-len(temp); i++ {
			temp2 += " "
		}
		fmt.Println(Yellow, "|", Violet, "                       Word to Find :", Green, temp, Yellow, temp2+"|")
		fmt.Println(Yellow, "|                                                                          |")
		temp = ""
		c := 0
		temp2 = ""
		for _, element := range AlreadyTry {
			temp += element + " "
			c += len(element) + 1
		}
		for i := 0; i < 35-c; i++ {
			temp2 += " "
		}
		fmt.Println(Yellow, "|", White, "                You already try :", Gray, temp, Yellow, temp2+"|")
		fmt.Println(Yellow, "|                                                                          |")
		if Lap == 10 {
			temp2 = ""
		} else {
			temp2 = " "
		}
		fmt.Println(Yellow, "|", Red, "                        Tries remaining :", Cyan, Lap, Yellow, temp2+"                         |")
		DisplayHangman(Lap)
		fmt.Println(Red, " ==========================================================================")
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
	DisplayFinish()
}
