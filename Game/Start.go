package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//ficher ou l'on met le debut du jeu

var Lap int
var Word string
var Lenght int

func Start(word string) {
	ClearScreen()
	Lap = 10
	fmt.Printf("============================== Hangman game =============================\n\nWelcome to the hangman game.\nWould you like a reminder of the rules ?\n(1) Yes   (2) No\n")
	Word = strings.ToLower(word)
	Lenght = len(word) - 1
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		DisplayRules()
	case "2":
		Game()
	default:
		Start(word)
	}
}

func DisplayRules() {
	ClearScreen()
	fmt.Println("============================== Hangman game =============================")
	fmt.Printf("Rules\n\t- 1 : ")
}

func ClearScreen() {
	var cmd *exec.Cmd
	// Détecter le système d'exploitation
	switch runtime.GOOS {
	case "windows":
		// Commande pour Windows
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		// Commande pour les systèmes Unix-like
		cmd = exec.Command("clear")
	}
	// Définir la sortie de la commande sur Stdout
	cmd.Stdout = os.Stdout
	cmd.Run()
}
