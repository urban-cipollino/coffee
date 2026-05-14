package main

import (
	"bufio"
	"coffee/internal/infrastructure/cli/command"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("usage")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "help":
			command.Help()
		case "menu":
			command.Menu()
		case "stock":
			err := command.Stock(parts)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "brew":
			err := command.Brew(parts)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "stats":
			command.Stats()
		default:
			fmt.Println("Unknown command:", parts[0])
		}
	}
}
