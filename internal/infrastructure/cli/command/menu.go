package command

import (
	"coffee/internal/infrastructure"
	"fmt"
)

func Menu() {
	for _, item := range infrastructure.Menu {
		fmt.Println(item.Name, item.Price)
	}
}
