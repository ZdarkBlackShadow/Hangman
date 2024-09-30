package game

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[91m"
	Green  = "\033[32m"
	Yellow = "\033[93m"
	Violet = "\033[95m"
	Cyan   = "\033[96m"
	White  = "\033[37m"
	Gray   = "\033[90m"
	Brown  = "\033[38;5;130m"
	Orange = "\033[38;5;208m"
)

func DisplayHangman(nb int) {
	//Function who display the progression of the hangman
	switch nb {
	case 10:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 9:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "O", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 8:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "O", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 7:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "O", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 6:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "O", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/|\\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 5:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "O", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/|\\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/ ", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 4:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "O", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/|\\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/ \\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 3:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "O", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |  ", White, "\\_|\\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/ \\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 2:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "O", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |  ", White, "\\_|_/", Yellow, "                              |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/ \\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	case 1:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "o", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |  ", White, "\\_|_/", Yellow, "                              |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/ \\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")
	default:
		fmt.Println(Yellow, "|                         ", Brown, "      _______", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |    ", White, "|", Yellow, "                                |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "[x]", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |  ", White, "\\_|_/", Yellow, "                              |")
		fmt.Println(Yellow, "|                         ", Brown, "     |   ", White, "/ \\", Yellow, "                               |")
		fmt.Println(Yellow, "|                         ", Brown, "     |", Yellow, "                                       |")
		fmt.Println(Yellow, "|                         ", Brown, "_____|_____", Yellow, "                                  |")

	}
}

func Formatage() {
	//Function to format with space the word to find
	Orange := "\033[31m"
	totalWidth := 74
	leftPadding := (totalWidth - len(Word)) / 2
	rightPadding := totalWidth - len(Word) - leftPadding
	fmt.Printf("%s |%s%s%s%s%s|%s\n", Yellow, strings.Repeat(" ", leftPadding), Orange, Word, Yellow, strings.Repeat(" ", rightPadding), Reset)
}

func FormatageDisplay(c string) string {
	res := Yellow + " |" + Gray
	l := len(c)
	n := (72 - l) / 2
	if l%2 == 1 {
		res += " "
	}
	for i := 0; i < n; i++ {
		res += " "
	}
	res += c
	for i := 0; i <= n+1; i++ {
		res += " "
	}
	return res + Yellow + "|"
}

func DisplayAlreadyTry(L []string) {
	fmt.Println(Yellow, "|                                                                          |")
	fmt.Println(Yellow, "|", White, "                         You already try :                            ", Yellow, "|")
	temp := 0
	ToDisplay := ""
	for _, element := range L {
		if element != "" {
			if temp+len(element) > 72 {
				fmt.Println(FormatageDisplay(ToDisplay))
				ToDisplay = ""
				temp = 0
			}
			temp += len(element) + 2
			if element == L[len(L)-1] {
				ToDisplay += element
			} else {
				ToDisplay += element + ", "
			}
		}
	}
	fmt.Println(FormatageDisplay(ToDisplay))
	fmt.Println(Yellow, "|                                                                          |")
}

func DisplayNbGameDifficulty() {
	fmt.Println(Yellow, "|                                                                          |")
	t := ""
	for i := 0; i < 41-len(strconv.Itoa(NbGame))-len(strconv.Itoa(Difficulte)); i++ {
		t += " "
	}
	fmt.Println(Yellow, "|", Green, "Game number : "+strconv.Itoa(NbGame)+t, Red, "Difficulty : "+strconv.Itoa(Difficulte), Yellow, "|")
	fmt.Println(Yellow, "|                                                                          |")
}
