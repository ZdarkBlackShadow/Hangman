package game

import (
	"fmt"
	"io/ioutil"
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

func DisplayHangman2(nb int) {
	content, error := ioutil.ReadFile("Data/Animation/Animation.txt")
	if error != nil {
		fmt.Println("error while opening animation.txt", error)
	}
	t := nb == 0
	if nb == 10 {
		nb = 0
	}
	temp := ""
	count := 7
	add := false
	for _, ele := range string(content) {
		if !add && ele == rune(48+nb) {
			if t {
				t = false
			} else {
				add = true
			}
		} else if add {
			switch ele {
			case 'Y':
				temp += Yellow
			case 'B':
				temp += Brown
			case 'W':
				temp += White
			case '\n':
				count--
				if count != 0 {
					temp += "\n"
				}
			default:
				temp += string(ele)
			}
		}
		if count <= 0 {
			break
		}
	}
	fmt.Println(temp)
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
