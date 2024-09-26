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
	RandomLetter := rune(Word[rand.Intn(Lenght)])
	ListToDisplay := []rune{}
	AlreadyTry := []string{}
	AlreadyTryCorrect := []rune{RandomLetter}
	LetterFind := 0
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
			SubmitesLetter := rune(choice[0])
			FindInWord := false
			for c := 0; c < Lenght; c++ {
				if SubmitesLetter == rune(Word[c]) && ListToDisplay[c] == '_' {
					ListToDisplay[c] = rune(SubmitesLetter)
					FindInWord = true
					LetterFind++
					AlreadyTryCorrect = append(AlreadyTryCorrect, rune(SubmitesLetter))
				}
			}
			for _, element := range AlreadyTryCorrect {
				if SubmitesLetter == element {
					FindInWord = true
				}
			}
			for _, element := range AlreadyTry {
				if len(element) == 1 {
					if SubmitesLetter == rune(element[0]) {
						FindInWord = true
					}
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
				IsAlreadyTry := false
				for _, element := range AlreadyTry {
					if len(element) != 1 && choice == element {
						IsAlreadyTry = true
					}
				}
				if !IsAlreadyTry {
					AlreadyTry = append(AlreadyTry, choice)
					Lap -= 2
				}
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
