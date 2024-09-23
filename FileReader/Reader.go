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
		fmt.Println("Erreur lors de l'ouverture du fucher", error)
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
