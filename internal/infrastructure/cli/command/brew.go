package command

import (
	"coffee/internal/domain"
	"coffee/internal/domain/entity"
	"coffee/internal/infrastructure"
	"fmt"
	"strconv"
)

func Brew(parts []string) error {
	if len(parts) < 4 {
		return domain.ErrInvalidParams
	}

	drinkName := parts[1]
	flag := parts[2]
	money, err := strconv.Atoi(parts[3])
	if err != nil {
		return domain.ErrInvalidParams
	}

	if flag != "--pay" {
		return domain.ErrInvalidParams
	}

	drink, ok := infrastructure.Menu[drinkName]
	if !ok {
		return domain.ErrDrinkNotFound
	}

	enoughMoney := drink.IsEnoughMoney(money)
	if !enoughMoney {
		return domain.ErrNotEnoughMoney
	}

	err = checkIngredients(drink)
	if err != nil {
		return err
	}

	err = decreaseIngredients(drink)
	if err != nil {
		return err
	}

	addStatistic(drink)

	for _, step := range drink.Recipe.Steps {
		fmt.Println(step)
	}

	return nil
}

func checkIngredients(drink entity.Drink) error {
	for _, ing := range drink.Recipe.Ingredients {
		stock := infrastructure.Ingredients[ing.Name]

		enough, err := stock.IsEnoughForDrink(drink)
		if err != nil {
			return err
		}
		if !enough {
			return domain.ErrNotEnoughIngredients
		}
	}

	return nil
}

func decreaseIngredients(drink entity.Drink) error {
	for _, ing := range drink.Recipe.Ingredients {
		stock := infrastructure.Ingredients[ing.Name]
		if err := stock.Decrease(ing.Quantity); err != nil {
			return err
		}

		infrastructure.Ingredients[ing.Name] = stock
	}

	return nil
}

func addStatistic(drink entity.Drink) {
	infrastructure.Statistic.OrderCount++
	infrastructure.Statistic.Revenue += drink.Price
}
