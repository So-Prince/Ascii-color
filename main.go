package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ChooseBanner(doc string, sub string, color string) {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := strings.Split(string(file), "\n")
	font := strings.Split(doc, "\\n")

	colorCode := GetColor(color)
	reset := "\033[0m"

	for _, ch := range font {
		if ch == "" {
			fmt.Println()
			continue
		}

		// 🔥 Step 1: mark positions to color
		mark := make([]bool, len(ch))

		if sub == "" {
			for i := range mark {
				mark[i] = true
			}
		} else {
			for i := 0; i <= len(ch)-len(sub); i++ {
				if ch[i:i+len(sub)] == sub {
					for j := 0; j < len(sub); j++ {
						mark[i+j] = true
					}
				}
			}
		}

		// 🔥 Step 2: print ASCII with color
		for line := 0; line < 8; line++ {
			for i, char := range ch {
				input := (char-32)*9 + rune(line)

				if mark[i] {
					fmt.Print(colorCode + result[input] + reset)
				} else {
					fmt.Print(result[input])
				}
			}
			fmt.Println()
		}
	}
}

func GetColor(color string) string {
	switch color {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	case "Magneta":
		return "\033[35m"
	default:
		return "\033[0m"
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}

	color := ""
	sub := ""
	text := ""

	// Case 1: with color flag
	if strings.HasPrefix(args[1], "--color=") {
		if len(args) < 3 {
			fmt.Println("Usage: go run . --color=<color> <substring> \"text\"")
			return
		}

		color = strings.TrimPrefix(args[1], "--color=")

		if len(args) == 3 {
			// no substring → color all
			text = args[2]
		} else {
			sub = args[2]
			text = args[3]
		}

	} else {
		// no color
		text = args[1]
	}

	ChooseBanner(text, sub, color)
}
