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
		fmt.Println("Error when opening file")
		fmt.Printf("\n\n")
		findN := false
		for i, element := range filename {
			switch element {
			case '1':
				if i+1 != len(filename) && filename[i+1] == '0' {
					fmt.Println("Try with this command : go run . File10.txt")
				} else {
					fmt.Println("Try with this command : go run . File1.txt")
				}
				findN = true
			case '2':
				fmt.Println("Try with this command : go run . File2File")
				findN = true
			case '3':
				fmt.Println("Try with this command : go run . File3.txt")
				findN = true
			case '4':
				fmt.Println("Try with this command : go run . File4.txt")
				findN = true
			case '5':
				fmt.Println("Try with this command : go run . File5.txt")
				findN = true
			case '6':
				fmt.Println("Try with this command : go run . File6.txt")
				findN = true
			case '7':
				fmt.Println("Try with this command : go run . File7.txt")
				findN = true
			case '8':
				fmt.Println("Try with this command : go run . File8.txt")
				findN = true
			case '9':
				fmt.Println("Try with this command : go run . File9.txt")
				findN = true
			case '0':
				fmt.Println("Try with this command : go run . File10.txt")
				findN = true
			}
			if findN {
				break
			}
		}
		if !findN {
			fmt.Println("Try with this command : go run . File1.txt")
		}
		fmt.Printf("\n\n")
	}
	temp := ""
	for _, element := range string(content) {
		if element != '\n' && element != '\r' {
			temp += string(element)
		} else {
			if temp != "" {
				res = append(res, temp)
				temp = ""
			}
		}
	}
	return res
}
