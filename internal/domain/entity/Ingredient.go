package entity

import (
	"coffee/internal/domain"
)

type Ingredient struct {
	Name     string
	Quantity int
}

func (i *Ingredient) Increase(quantity int) error {
	if quantity < 1 {
		return domain.ErrInvalidParams
	}

	i.Quantity += quantity

	return nil
}

func (i *Ingredient) Decrease(quantity int) error {
	if i.Quantity-quantity < 0 {
		return domain.ErrInvalidParams
	}

	i.Quantity -= quantity

	return nil
}

func (i *Ingredient) Set(quantity int) error {
	if quantity < 0 {
		return domain.ErrInvalidParams
	}

	i.Quantity = quantity

	return nil
}

func (i *Ingredient) IsEnoughForDrink(drink Drink) (bool, error) {
	for _, ing := range drink.Recipe.Ingredients {
		if i.Name == ing.Name {
			return i.Quantity >= ing.Quantity, nil
		}
	}

	return false, domain.ErrIngredientIsNotUsedForDrink
}
