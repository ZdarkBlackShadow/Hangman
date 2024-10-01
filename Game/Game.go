package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var Finish bool //variable who indicate if the game if finish or not.
var Win bool    //variable to indicate if the user win or not

func Game() {
	/*
		Function who are in charge of the game.
	*/
	rand.Seed(time.Now().UnixNano())
	Tries = 12
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
		DisplayNbGameDifficulty()
		temp := ""
		for _, element := range ListToDisplay {
			temp += string(element) + " "
		}
		temp2 := ""
		for i := 0; i < 31-len(temp); i++ {
			temp2 += " "
		}
		fmt.Println(Yellow, "|", Violet, "                       Word to Find :", Green, temp, Yellow, temp2+"|")
		DisplayAlreadyTry(AlreadyTry)
		if Tries >= 10 {
			temp2 = ""
		} else {
			temp2 = " "
		}
		fmt.Println(Yellow, "|", Red, "                        Tries remaining :", Cyan, Tries, Yellow, temp2+"                         |")
		DisplayHangman2(Tries)
		fmt.Println(Red, " ==========================================================================")
		fmt.Printf("%s\n\nEnter a letter or a word : ", White)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		temp = scanner.Text()
		temp = strings.ToLower(temp)
		temp = strings.ReplaceAll(temp, " ", "")
		if verif(temp) {
			if len(temp) == 1 {
				SubmitesLetter := rune(temp[0])
				if SubmitesLetter >= 'a' && SubmitesLetter <= 'z' {
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
						AlreadyTry = append(AlreadyTry, temp)
						Tries--
					}
				}
			} else {
				if temp == Word {
					LetterFind = Lenght
				} else {
					IsAlreadyTry := false
					for _, element := range AlreadyTry {
						if len(element) != 1 && temp == element {
							IsAlreadyTry = true
						}
					}
					if !IsAlreadyTry {
						AlreadyTry = append(AlreadyTry, temp)
						Tries -= 2
					}
				}
			}
		}
		if LetterFind == Lenght {
			Finish = true
			Win = true
		} else {
			if Tries <= 0 {
				Finish = true
			}
		}
	}
}

func verif(answer string) bool {
	for _, element := range answer {
		if !(element >= 'a' && element <= 'z') {
			return false
		}
	}
	return true
}
