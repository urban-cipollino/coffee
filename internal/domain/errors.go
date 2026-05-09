package domain

import "errors"

var (
	ErrDrinkNotFound               = errors.New("the drink was not found")
	ErrIngredientIsNotUsedForDrink = errors.New("the ingredient is not used for the drink")
	ErrNotEnoughMoney              = errors.New("not enough money")
	ErrNotEnoughIngredients        = errors.New("not enough ingredients")
	ErrInvalidParams               = errors.New("invalid params")
)
