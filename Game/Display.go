package game

import "fmt"

func DisplayHangman(nb int) {
	switch nb {
	case 10:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 9:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     O")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 8:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     O")
		fmt.Println("     |     |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 7:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     O")
		fmt.Println("     |    /|")
		fmt.Println("     |      ")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 6:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     O")
		fmt.Println("     |    /|\\")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 5:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     O")
		fmt.Println("     |    /|\\")
		fmt.Println("     |    /")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 4:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     O")
		fmt.Println("     |    /|\\")
		fmt.Println("     |    / \\")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 3:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     O")
		fmt.Println("     |   \\_|/")
		fmt.Println("     |    / \\")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 2:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     O")
		fmt.Println("     |   \\_|_/")
		fmt.Println("     |    / \\")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	case 1:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |     o")
		fmt.Println("     |   \\_|_/")
		fmt.Println("     |    / \\")
		fmt.Println("     |")
		fmt.Println("_____|_____")
	default:
		fmt.Println("      _______")
		fmt.Println("     |     |")
		fmt.Println("     |    [x]")
		fmt.Println("     |   \\_|_/")
		fmt.Println("     |    / \\")
		fmt.Println("     |")
		fmt.Println("_____|_____")

	}
}
