package command

import (
	"coffee/internal/domain"
	"coffee/internal/infrastructure"
	"fmt"
	"strconv"
)

func Stock(parts []string) error {
	if len(parts) < 2 {
		return domain.ErrInvalidParams
	}

	commandName := parts[1]

	switch commandName {
	case "add", "set":
		if len(parts) < 4 {
			return domain.ErrInvalidParams
		}

		ingredientName := parts[2]
		quantity, err := strconv.Atoi(parts[3])
		if err != nil {
			return domain.ErrInvalidParams
		}

		ingredient, ok := infrastructure.Ingredients[ingredientName]
		if !ok {
			return domain.ErrInvalidParams
		}

		switch commandName {
		case "add":
			if err = ingredient.Increase(quantity); err != nil {
				return err
			}
		case "set":
			if err = ingredient.Set(quantity); err != nil {
				return err
			}
		}

		infrastructure.Ingredients[ingredientName] = ingredient

		fmt.Println("ok")
	case "get":
		for _, ing := range infrastructure.Ingredients {
			fmt.Println(ing.Name, ing.Quantity)
		}
	default:
		return domain.ErrInvalidParams
	}

	return nil
}
