package game

import "fmt"

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

func Formatage(c string, z int) string {
	res := ""
	l := len(c)
	n := (z - l) / 2
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
	return res
}
