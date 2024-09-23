package filereader

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// ficher ou l'on lit les ficher txt
func RandomWord(filename string) string {
	rand.Seed(time.Now().UnixNano())
	L := Reader(filename)
	return L[rand.Intn(len(L))]
}

func Reader(filename string) []string {
	res := []string{}
	content, error := ioutil.ReadFile("Data\\" + filename)
	if error != nil {
		fmt.Println("Error when opening file", error)
		fmt.Printf("\n\n\n")
		findN := false
		for i, element := range filename {
			switch element {
			case '1':
				if i+1 != len(filename) && filename[i+1] == '0' {
					fmt.Println("Try with this command : go run . Ficher10.txt")
				} else {
					fmt.Println("Try with this command : go run . Ficher1.txt")
				}
				findN = true
			case '2':
				fmt.Println("Try with this command : go run . Ficher2.txt")
				findN = true
			case '3':
				fmt.Println("Try with this command : go run . Ficher3.txt")
				findN = true
			case '4':
				fmt.Println("Try with this command : go run . Ficher4.txt")
				findN = true
			case '5':
				fmt.Println("Try with this command : go run . Ficher5.txt")
				findN = true
			case '6':
				fmt.Println("Try with this command : go run . Ficher6.txt")
				findN = true
			case '7':
				fmt.Println("Try with this command : go run . Ficher7.txt")
				findN = true
			case '8':
				fmt.Println("Try with this command : go run . Ficher8.txt")
				findN = true
			case '9':
				fmt.Println("Try with this command : go run . Ficher9.txt")
				findN = true
			case '0':
				fmt.Println("Try with this command : go run . Ficher10.txt")
				findN = true
			}
			if findN {
				break
			}
		}
		if !findN {
			fmt.Println("Try with this command : go run . Ficher1.txt")
		}
		fmt.Printf("\n\n\n\n")
	}
	temp := ""
	for _, element := range string(content) {
		if element != '\n' {
			temp += string(element)
		} else {
			res = append(res, temp)
			temp = ""
		}
	}
	return res
}
