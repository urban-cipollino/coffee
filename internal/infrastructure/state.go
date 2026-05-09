package infrastructure

import "coffee/internal/domain/entity"

var Statistic = struct {
	OrderCount, Revenue int
}{}

var Ingredients = map[string]entity.Ingredient{
	"beans": {Name: "beans", Quantity: 0},
	"water": {Name: "water", Quantity: 0},
	"milk":  {Name: "milk", Quantity: 0},
}

var Menu = map[string]entity.Drink{
	"espresso": {
		Name:  "Espresso",
		Price: 150,
		Recipe: entity.Recipe{
			Ingredients: []entity.Ingredient{
				{Name: "beans", Quantity: 8},
				{Name: "water", Quantity: 30},
			},
			Steps: []string{"grind", "tamp", "brew 25s"},
		},
	},
	"americano": {
		Name:  "Americano",
		Price: 180,
		Recipe: entity.Recipe{
			Ingredients: []entity.Ingredient{
				{Name: "beans", Quantity: 8},
				{Name: "water", Quantity: 120},
			},
			Steps: []string{"grind", "brew 25s", "add water"},
		},
	},
	"latte": {
		Name:  "Latte",
		Price: 220,
		Recipe: entity.Recipe{
			Ingredients: []entity.Ingredient{
				{Name: "beans", Quantity: 8},
				{Name: "water", Quantity: 30},
				{Name: "milk", Quantity: 150},
			},
			Steps: []string{"grind", "brew 25s", "steam milk", "mix"},
		},
	},
}
